"""An addon using the abbreviated scripting syntax."""
import uuid

def request(flow):
    flow.request.headers["dt-mark-header"] = uuid.uuid4().hex
