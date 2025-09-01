# vkResetEvent(3) Manual Page

## Name

vkResetEvent - Reset an event to non-signaled state



## [](#_c_specification)C Specification

To set the state of an event to unsignaled from the host, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkResetEvent(
    VkDevice                                    device,
    VkEvent                                     event);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the event.
- `event` is the event to reset.

## [](#_description)Description

When [vkResetEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/vkResetEvent.html) is executed on the host, it defines an *event unsignal operation* which resets the event to the unsignaled state.

If `event` is already in the unsignaled state when [vkResetEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/vkResetEvent.html) is executed, then [vkResetEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/vkResetEvent.html) has no effect, and no event unsignal operation occurs.

Valid Usage

- [](#VUID-vkResetEvent-event-03821)VUID-vkResetEvent-event-03821  
  There **must** be an execution dependency between `vkResetEvent` and the execution of any [vkCmdWaitEvents](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWaitEvents.html) that includes `event` in its `pEvents` parameter
- [](#VUID-vkResetEvent-event-03822)VUID-vkResetEvent-event-03822  
  There **must** be an execution dependency between `vkResetEvent` and the execution of any [vkCmdWaitEvents2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWaitEvents2.html) that includes `event` in its `pEvents` parameter
- [](#VUID-vkResetEvent-event-03823)VUID-vkResetEvent-event-03823  
  `event` **must** not have been created with `VK_EVENT_CREATE_DEVICE_ONLY_BIT`

Valid Usage (Implicit)

- [](#VUID-vkResetEvent-device-parameter)VUID-vkResetEvent-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkResetEvent-event-parameter)VUID-vkResetEvent-event-parameter  
  `event` **must** be a valid [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html) handle
- [](#VUID-vkResetEvent-event-parent)VUID-vkResetEvent-event-parent  
  `event` **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `event` **must** be externally synchronized

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkResetEvent).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0