# vkResetEvent(3) Manual Page

## Name

vkResetEvent - Reset an event to non-signaled state



## <a href="#_c_specification" class="anchor"></a>C Specification

To set the state of an event to unsignaled from the host, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkResetEvent(
    VkDevice                                    device,
    VkEvent                                     event);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the event.

- `event` is the event to reset.

## <a href="#_description" class="anchor"></a>Description

When [vkResetEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkResetEvent.html) is executed on the host, it
defines an *event unsignal operation* which resets the event to the
unsignaled state.

If `event` is already in the unsignaled state when
[vkResetEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkResetEvent.html) is executed, then
[vkResetEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkResetEvent.html) has no effect, and no event unsignal
operation occurs.

Valid Usage

- <a href="#VUID-vkResetEvent-event-03821"
  id="VUID-vkResetEvent-event-03821"></a>
  VUID-vkResetEvent-event-03821  
  There **must** be an execution dependency between `vkResetEvent` and
  the execution of any [vkCmdWaitEvents](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents.html) that
  includes `event` in its `pEvents` parameter

- <a href="#VUID-vkResetEvent-event-03822"
  id="VUID-vkResetEvent-event-03822"></a>
  VUID-vkResetEvent-event-03822  
  There **must** be an execution dependency between `vkResetEvent` and
  the execution of any [vkCmdWaitEvents2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents2.html) that
  includes `event` in its `pEvents` parameter

- <a href="#VUID-vkResetEvent-event-03823"
  id="VUID-vkResetEvent-event-03823"></a>
  VUID-vkResetEvent-event-03823  
  `event` **must** not have been created with
  `VK_EVENT_CREATE_DEVICE_ONLY_BIT`

Valid Usage (Implicit)

- <a href="#VUID-vkResetEvent-device-parameter"
  id="VUID-vkResetEvent-device-parameter"></a>
  VUID-vkResetEvent-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkResetEvent-event-parameter"
  id="VUID-vkResetEvent-event-parameter"></a>
  VUID-vkResetEvent-event-parameter  
  `event` **must** be a valid [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html) handle

- <a href="#VUID-vkResetEvent-event-parent"
  id="VUID-vkResetEvent-event-parent"></a>
  VUID-vkResetEvent-event-parent  
  `event` **must** have been created, allocated, or retrieved from
  `device`

Host Synchronization

- Host access to `event` **must** be externally synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkResetEvent"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
