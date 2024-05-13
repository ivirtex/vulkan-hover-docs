# vkDestroyVideoSessionKHR(3) Manual Page

## Name

vkDestroyVideoSessionKHR - Destroy video session object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a video session, call:

``` c
// Provided by VK_KHR_video_queue
void vkDestroyVideoSessionKHR(
    VkDevice                                    device,
    VkVideoSessionKHR                           videoSession,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the video session.

- `videoSession` is the video session to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkDestroyVideoSessionKHR-videoSession-07192"
  id="VUID-vkDestroyVideoSessionKHR-videoSession-07192"></a>
  VUID-vkDestroyVideoSessionKHR-videoSession-07192  
  All submitted commands that refer to `videoSession` **must** have
  completed execution

- <a href="#VUID-vkDestroyVideoSessionKHR-videoSession-07193"
  id="VUID-vkDestroyVideoSessionKHR-videoSession-07193"></a>
  VUID-vkDestroyVideoSessionKHR-videoSession-07193  
  If `VkAllocationCallbacks` were provided when `videoSession` was
  created, a compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroyVideoSessionKHR-videoSession-07194"
  id="VUID-vkDestroyVideoSessionKHR-videoSession-07194"></a>
  VUID-vkDestroyVideoSessionKHR-videoSession-07194  
  If no `VkAllocationCallbacks` were provided when `videoSession` was
  created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyVideoSessionKHR-device-parameter"
  id="VUID-vkDestroyVideoSessionKHR-device-parameter"></a>
  VUID-vkDestroyVideoSessionKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyVideoSessionKHR-videoSession-parameter"
  id="VUID-vkDestroyVideoSessionKHR-videoSession-parameter"></a>
  VUID-vkDestroyVideoSessionKHR-videoSession-parameter  
  If `videoSession` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `videoSession` **must** be a valid
  [VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionKHR.html) handle

- <a href="#VUID-vkDestroyVideoSessionKHR-pAllocator-parameter"
  id="VUID-vkDestroyVideoSessionKHR-pAllocator-parameter"></a>
  VUID-vkDestroyVideoSessionKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyVideoSessionKHR-videoSession-parent"
  id="VUID-vkDestroyVideoSessionKHR-videoSession-parent"></a>
  VUID-vkDestroyVideoSessionKHR-videoSession-parent  
  If `videoSession` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `videoSession` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyVideoSessionKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
