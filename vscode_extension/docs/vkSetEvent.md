# vkSetEvent(3) Manual Page

## Name

vkSetEvent - Set an event to signaled state



## [](#_c_specification)C Specification

To set the state of an event to signaled from the host, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkSetEvent(
    VkDevice                                    device,
    VkEvent                                     event);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the event.
- `event` is the event to set.

## [](#_description)Description

When [vkSetEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetEvent.html) is executed on the host, it defines an *event signal operation* which sets the event to the signaled state.

If `event` is already in the signaled state when [vkSetEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetEvent.html) is executed, then [vkSetEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetEvent.html) has no effect, and no event signal operation occurs.

Note

If a command buffer is waiting for an event to be signaled from the host, the application must signal the event before submitting the command buffer, as described in the [queue forward progress](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-submission-progress) section.

Valid Usage

- [](#VUID-vkSetEvent-event-03941)VUID-vkSetEvent-event-03941  
  `event` **must** not have been created with `VK_EVENT_CREATE_DEVICE_ONLY_BIT`
- [](#VUID-vkSetEvent-event-09543)VUID-vkSetEvent-event-09543  
  `event` **must** not be waited on by a command buffer in the [pending state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle)

Valid Usage (Implicit)

- [](#VUID-vkSetEvent-device-parameter)VUID-vkSetEvent-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkSetEvent-event-parameter)VUID-vkSetEvent-event-parameter  
  `event` **must** be a valid [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html) handle
- [](#VUID-vkSetEvent-event-parent)VUID-vkSetEvent-event-parent  
  `event` **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `event` **must** be externally synchronized

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkSetEvent).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0