# vkGetEventStatus(3) Manual Page

## Name

vkGetEventStatus - Retrieve the status of an event object



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the state of an event from the host, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkGetEventStatus(
    VkDevice                                    device,
    VkEvent                                     event);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the event.

- `event` is the handle of the event to query.

## <a href="#_description" class="anchor"></a>Description

Upon success, `vkGetEventStatus` returns the state of the event object
with the following return codes:

| Status           | Meaning                                       |
|------------------|-----------------------------------------------|
| `VK_EVENT_SET`   | The event specified by `event` is signaled.   |
| `VK_EVENT_RESET` | The event specified by `event` is unsignaled. |

Table 1. Event Object Status Codes

If a `vkCmdSetEvent` or `vkCmdResetEvent` command is in a command buffer
that is in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
target="_blank" rel="noopener">pending state</a>, then the value
returned by this command **may** immediately be out of date.

The state of an event **can** be updated by the host. The state of the
event is immediately changed, and subsequent calls to `vkGetEventStatus`
will return the new state. If an event is already in the requested
state, then updating it to the same state has no effect.

Valid Usage

- <a href="#VUID-vkGetEventStatus-event-03940"
  id="VUID-vkGetEventStatus-event-03940"></a>
  VUID-vkGetEventStatus-event-03940  
  `event` **must** not have been created with
  `VK_EVENT_CREATE_DEVICE_ONLY_BIT`

Valid Usage (Implicit)

- <a href="#VUID-vkGetEventStatus-device-parameter"
  id="VUID-vkGetEventStatus-device-parameter"></a>
  VUID-vkGetEventStatus-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetEventStatus-event-parameter"
  id="VUID-vkGetEventStatus-event-parameter"></a>
  VUID-vkGetEventStatus-event-parameter  
  `event` **must** be a valid [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html) handle

- <a href="#VUID-vkGetEventStatus-event-parent"
  id="VUID-vkGetEventStatus-event-parent"></a>
  VUID-vkGetEventStatus-event-parent  
  `event` **must** have been created, allocated, or retrieved from
  `device`

Return Codes

On success, this command returns  
- `VK_EVENT_SET`

- `VK_EVENT_RESET`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_DEVICE_LOST`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetEventStatus"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
