# vkDestroyEvent(3) Manual Page

## Name

vkDestroyEvent - Destroy an event object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy an event, call:

``` c
// Provided by VK_VERSION_1_0
void vkDestroyEvent(
    VkDevice                                    device,
    VkEvent                                     event,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the event.

- `event` is the handle of the event to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkDestroyEvent-event-01145"
  id="VUID-vkDestroyEvent-event-01145"></a>
  VUID-vkDestroyEvent-event-01145  
  All submitted commands that refer to `event` **must** have completed
  execution

- <a href="#VUID-vkDestroyEvent-event-01146"
  id="VUID-vkDestroyEvent-event-01146"></a>
  VUID-vkDestroyEvent-event-01146  
  If `VkAllocationCallbacks` were provided when `event` was created, a
  compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroyEvent-event-01147"
  id="VUID-vkDestroyEvent-event-01147"></a>
  VUID-vkDestroyEvent-event-01147  
  If no `VkAllocationCallbacks` were provided when `event` was created,
  `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyEvent-device-parameter"
  id="VUID-vkDestroyEvent-device-parameter"></a>
  VUID-vkDestroyEvent-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyEvent-event-parameter"
  id="VUID-vkDestroyEvent-event-parameter"></a>
  VUID-vkDestroyEvent-event-parameter  
  If `event` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `event`
  **must** be a valid [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html) handle

- <a href="#VUID-vkDestroyEvent-pAllocator-parameter"
  id="VUID-vkDestroyEvent-pAllocator-parameter"></a>
  VUID-vkDestroyEvent-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyEvent-event-parent"
  id="VUID-vkDestroyEvent-event-parent"></a>
  VUID-vkDestroyEvent-event-parent  
  If `event` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `event` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyEvent"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
