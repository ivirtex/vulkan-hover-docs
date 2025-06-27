# vkDestroyBuffer(3) Manual Page

## Name

vkDestroyBuffer - Destroy a buffer object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a buffer, call:

``` c
// Provided by VK_VERSION_1_0
void vkDestroyBuffer(
    VkDevice                                    device,
    VkBuffer                                    buffer,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the buffer.

- `buffer` is the buffer to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkDestroyBuffer-buffer-00922"
  id="VUID-vkDestroyBuffer-buffer-00922"></a>
  VUID-vkDestroyBuffer-buffer-00922  
  All submitted commands that refer to `buffer`, either directly or via
  a `VkBufferView`, **must** have completed execution

- <a href="#VUID-vkDestroyBuffer-buffer-00923"
  id="VUID-vkDestroyBuffer-buffer-00923"></a>
  VUID-vkDestroyBuffer-buffer-00923  
  If `VkAllocationCallbacks` were provided when `buffer` was created, a
  compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroyBuffer-buffer-00924"
  id="VUID-vkDestroyBuffer-buffer-00924"></a>
  VUID-vkDestroyBuffer-buffer-00924  
  If no `VkAllocationCallbacks` were provided when `buffer` was created,
  `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyBuffer-device-parameter"
  id="VUID-vkDestroyBuffer-device-parameter"></a>
  VUID-vkDestroyBuffer-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyBuffer-buffer-parameter"
  id="VUID-vkDestroyBuffer-buffer-parameter"></a>
  VUID-vkDestroyBuffer-buffer-parameter  
  If `buffer` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `buffer`
  **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-vkDestroyBuffer-pAllocator-parameter"
  id="VUID-vkDestroyBuffer-pAllocator-parameter"></a>
  VUID-vkDestroyBuffer-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyBuffer-buffer-parent"
  id="VUID-vkDestroyBuffer-buffer-parent"></a>
  VUID-vkDestroyBuffer-buffer-parent  
  If `buffer` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `buffer` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyBuffer"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
