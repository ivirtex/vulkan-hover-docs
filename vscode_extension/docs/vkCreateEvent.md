# vkCreateEvent(3) Manual Page

## Name

vkCreateEvent - Create a new event object



## [](#_c_specification)C Specification

To create an event, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkCreateEvent(
    VkDevice                                    device,
    const VkEventCreateInfo*                    pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkEvent*                                    pEvent);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the event.
- `pCreateInfo` is a pointer to a [VkEventCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEventCreateInfo.html) structure containing information about how the event is to be created.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pEvent` is a pointer to a handle in which the resulting event object is returned.

## [](#_description)Description

When created, the event object is in the unsignaled state.

Valid Usage

- [](#VUID-vkCreateEvent-device-09672)VUID-vkCreateEvent-device-09672  
  `device` **must** support at least one queue family with one of the `VK_QUEUE_VIDEO_ENCODE_BIT_KHR`, `VK_QUEUE_VIDEO_DECODE_BIT_KHR`, `VK_QUEUE_COMPUTE_BIT`, or `VK_QUEUE_GRAPHICS_BIT` capabilities
- [](#VUID-vkCreateEvent-events-04468)VUID-vkCreateEvent-events-04468  
  If the `VK_KHR_portability_subset` extension is enabled, and [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`events` is `VK_FALSE`, then the implementation does not support [events](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-events), and [vkCreateEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateEvent.html) **must** not be used

Valid Usage (Implicit)

- [](#VUID-vkCreateEvent-device-parameter)VUID-vkCreateEvent-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateEvent-pCreateInfo-parameter)VUID-vkCreateEvent-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkEventCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEventCreateInfo.html) structure
- [](#VUID-vkCreateEvent-pAllocator-parameter)VUID-vkCreateEvent-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateEvent-pEvent-parameter)VUID-vkCreateEvent-pEvent-parameter  
  `pEvent` **must** be a valid pointer to a [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html) handle
- [](#VUID-vkCreateEvent-device-queuecount)VUID-vkCreateEvent-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html), [VkEventCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEventCreateInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateEvent).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0