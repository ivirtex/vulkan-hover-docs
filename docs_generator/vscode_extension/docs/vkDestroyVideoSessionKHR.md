# vkDestroyVideoSessionKHR(3) Manual Page

## Name

vkDestroyVideoSessionKHR - Destroy video session object



## [](#_c_specification)C Specification

To destroy a video session, call:

```c++
// Provided by VK_KHR_video_queue
void vkDestroyVideoSessionKHR(
    VkDevice                                    device,
    VkVideoSessionKHR                           videoSession,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the video session.
- `videoSession` is the video session to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyVideoSessionKHR-videoSession-07192)VUID-vkDestroyVideoSessionKHR-videoSession-07192  
  All submitted commands that refer to `videoSession` **must** have completed execution
- [](#VUID-vkDestroyVideoSessionKHR-videoSession-07193)VUID-vkDestroyVideoSessionKHR-videoSession-07193  
  If `VkAllocationCallbacks` were provided when `videoSession` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyVideoSessionKHR-videoSession-07194)VUID-vkDestroyVideoSessionKHR-videoSession-07194  
  If no `VkAllocationCallbacks` were provided when `videoSession` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroyVideoSessionKHR-device-parameter)VUID-vkDestroyVideoSessionKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyVideoSessionKHR-videoSession-parameter)VUID-vkDestroyVideoSessionKHR-videoSession-parameter  
  If `videoSession` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `videoSession` **must** be a valid [VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionKHR.html) handle
- [](#VUID-vkDestroyVideoSessionKHR-pAllocator-parameter)VUID-vkDestroyVideoSessionKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyVideoSessionKHR-videoSession-parent)VUID-vkDestroyVideoSessionKHR-videoSession-parent  
  If `videoSession` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `videoSession` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyVideoSessionKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0