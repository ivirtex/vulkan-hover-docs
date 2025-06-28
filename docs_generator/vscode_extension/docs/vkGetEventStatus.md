# vkGetEventStatus(3) Manual Page

## Name

vkGetEventStatus - Retrieve the status of an event object



## [](#_c_specification)C Specification

To query the state of an event from the host, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkGetEventStatus(
    VkDevice                                    device,
    VkEvent                                     event);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the event.
- `event` is the handle of the event to query.

## [](#_description)Description

Upon success, `vkGetEventStatus` returns the state of the event object with the following return codes:

Table 1. Event Object Status Codes   Status Meaning

`VK_EVENT_SET`

The event specified by `event` is signaled.

`VK_EVENT_RESET`

The event specified by `event` is unsignaled.

If a `vkCmdSetEvent` or `vkCmdResetEvent` command is in a command buffer that is in the [pending state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle), then the value returned by this command **may** immediately be out of date.

The state of an event **can** be updated by the host. The state of the event is immediately changed, and subsequent calls to `vkGetEventStatus` will return the new state. If an event is already in the requested state, then updating it to the same state has no effect.

Valid Usage

- [](#VUID-vkGetEventStatus-event-03940)VUID-vkGetEventStatus-event-03940  
  `event` **must** not have been created with `VK_EVENT_CREATE_DEVICE_ONLY_BIT`

Valid Usage (Implicit)

- [](#VUID-vkGetEventStatus-device-parameter)VUID-vkGetEventStatus-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetEventStatus-event-parameter)VUID-vkGetEventStatus-event-parameter  
  `event` **must** be a valid [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html) handle
- [](#VUID-vkGetEventStatus-event-parent)VUID-vkGetEventStatus-event-parent  
  `event` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_EVENT_SET`
- `VK_EVENT_RESET`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_DEVICE_LOST`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetEventStatus)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0