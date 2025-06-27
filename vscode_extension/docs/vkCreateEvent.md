# vkCreateEvent(3) Manual Page

## Name

vkCreateEvent - Create a new event object



## <a href="#_c_specification" class="anchor"></a>C Specification

To create an event, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkCreateEvent(
    VkDevice                                    device,
    const VkEventCreateInfo*                    pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkEvent*                                    pEvent);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the event.

- `pCreateInfo` is a pointer to a
  [VkEventCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEventCreateInfo.html) structure containing
  information about how the event is to be created.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pEvent` is a pointer to a handle in which the resulting event object
  is returned.

## <a href="#_description" class="anchor"></a>Description

When created, the event object is in the unsignaled state.

Valid Usage

- <a href="#VUID-vkCreateEvent-device-09672"
  id="VUID-vkCreateEvent-device-09672"></a>
  VUID-vkCreateEvent-device-09672  
  `device` **must** support at least one queue family with one of the
  `VK_QUEUE_VIDEO_ENCODE_BIT_KHR`, `VK_QUEUE_VIDEO_DECODE_BIT_KHR`,
  `VK_QUEUE_COMPUTE_BIT`, or `VK_QUEUE_GRAPHICS_BIT` capabilities

- <a href="#VUID-vkCreateEvent-events-04468"
  id="VUID-vkCreateEvent-events-04468"></a>
  VUID-vkCreateEvent-events-04468  
  If the [`VK_KHR_portability_subset`](VK_KHR_portability_subset.html)
  extension is enabled, and
  [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`events`
  is `VK_FALSE`, then the implementation does not support <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-events"
  target="_blank" rel="noopener">events</a>, and
  [vkCreateEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateEvent.html) **must** not be used

Valid Usage (Implicit)

- <a href="#VUID-vkCreateEvent-device-parameter"
  id="VUID-vkCreateEvent-device-parameter"></a>
  VUID-vkCreateEvent-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateEvent-pCreateInfo-parameter"
  id="VUID-vkCreateEvent-pCreateInfo-parameter"></a>
  VUID-vkCreateEvent-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkEventCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEventCreateInfo.html) structure

- <a href="#VUID-vkCreateEvent-pAllocator-parameter"
  id="VUID-vkCreateEvent-pAllocator-parameter"></a>
  VUID-vkCreateEvent-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateEvent-pEvent-parameter"
  id="VUID-vkCreateEvent-pEvent-parameter"></a>
  VUID-vkCreateEvent-pEvent-parameter  
  `pEvent` **must** be a valid pointer to a [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html)
  handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html),
[VkEventCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEventCreateInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateEvent"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
