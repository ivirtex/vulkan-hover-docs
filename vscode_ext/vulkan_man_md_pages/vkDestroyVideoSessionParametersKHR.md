# vkDestroyVideoSessionParametersKHR(3) Manual Page

## Name

vkDestroyVideoSessionParametersKHR - Destroy video session parameters
object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a video session parameters object, call:

``` c
// Provided by VK_KHR_video_queue
void vkDestroyVideoSessionParametersKHR(
    VkDevice                                    device,
    VkVideoSessionParametersKHR                 videoSessionParameters,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the video session
  parameters object.

- `videoSessionParameters` is the video session parameters object to
  destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-vkDestroyVideoSessionParametersKHR-videoSessionParameters-07212"
  id="VUID-vkDestroyVideoSessionParametersKHR-videoSessionParameters-07212"></a>
  VUID-vkDestroyVideoSessionParametersKHR-videoSessionParameters-07212  
  All submitted commands that refer to `videoSessionParameters` **must**
  have completed execution

- <a
  href="#VUID-vkDestroyVideoSessionParametersKHR-videoSessionParameters-07213"
  id="VUID-vkDestroyVideoSessionParametersKHR-videoSessionParameters-07213"></a>
  VUID-vkDestroyVideoSessionParametersKHR-videoSessionParameters-07213  
  If `VkAllocationCallbacks` were provided when `videoSessionParameters`
  was created, a compatible set of callbacks **must** be provided here

- <a
  href="#VUID-vkDestroyVideoSessionParametersKHR-videoSessionParameters-07214"
  id="VUID-vkDestroyVideoSessionParametersKHR-videoSessionParameters-07214"></a>
  VUID-vkDestroyVideoSessionParametersKHR-videoSessionParameters-07214  
  If no `VkAllocationCallbacks` were provided when
  `videoSessionParameters` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyVideoSessionParametersKHR-device-parameter"
  id="VUID-vkDestroyVideoSessionParametersKHR-device-parameter"></a>
  VUID-vkDestroyVideoSessionParametersKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkDestroyVideoSessionParametersKHR-videoSessionParameters-parameter"
  id="VUID-vkDestroyVideoSessionParametersKHR-videoSessionParameters-parameter"></a>
  VUID-vkDestroyVideoSessionParametersKHR-videoSessionParameters-parameter  
  If `videoSessionParameters` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `videoSessionParameters`
  **must** be a valid
  [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersKHR.html) handle

- <a href="#VUID-vkDestroyVideoSessionParametersKHR-pAllocator-parameter"
  id="VUID-vkDestroyVideoSessionParametersKHR-pAllocator-parameter"></a>
  VUID-vkDestroyVideoSessionParametersKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a
  href="#VUID-vkDestroyVideoSessionParametersKHR-videoSessionParameters-parent"
  id="VUID-vkDestroyVideoSessionParametersKHR-videoSessionParameters-parent"></a>
  VUID-vkDestroyVideoSessionParametersKHR-videoSessionParameters-parent  
  If `videoSessionParameters` is a valid handle, it **must** have been
  created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `videoSessionParameters` **must** be externally
  synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyVideoSessionParametersKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
