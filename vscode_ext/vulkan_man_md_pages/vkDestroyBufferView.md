# vkDestroyBufferView(3) Manual Page

## Name

vkDestroyBufferView - Destroy a buffer view object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a buffer view, call:

``` c
// Provided by VK_VERSION_1_0
void vkDestroyBufferView(
    VkDevice                                    device,
    VkBufferView                                bufferView,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the buffer view.

- `bufferView` is the buffer view to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkDestroyBufferView-bufferView-00936"
  id="VUID-vkDestroyBufferView-bufferView-00936"></a>
  VUID-vkDestroyBufferView-bufferView-00936  
  All submitted commands that refer to `bufferView` **must** have
  completed execution

- <a href="#VUID-vkDestroyBufferView-bufferView-00937"
  id="VUID-vkDestroyBufferView-bufferView-00937"></a>
  VUID-vkDestroyBufferView-bufferView-00937  
  If `VkAllocationCallbacks` were provided when `bufferView` was
  created, a compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroyBufferView-bufferView-00938"
  id="VUID-vkDestroyBufferView-bufferView-00938"></a>
  VUID-vkDestroyBufferView-bufferView-00938  
  If no `VkAllocationCallbacks` were provided when `bufferView` was
  created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyBufferView-device-parameter"
  id="VUID-vkDestroyBufferView-device-parameter"></a>
  VUID-vkDestroyBufferView-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyBufferView-bufferView-parameter"
  id="VUID-vkDestroyBufferView-bufferView-parameter"></a>
  VUID-vkDestroyBufferView-bufferView-parameter  
  If `bufferView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `bufferView` **must** be a valid [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html)
  handle

- <a href="#VUID-vkDestroyBufferView-pAllocator-parameter"
  id="VUID-vkDestroyBufferView-pAllocator-parameter"></a>
  VUID-vkDestroyBufferView-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyBufferView-bufferView-parent"
  id="VUID-vkDestroyBufferView-bufferView-parent"></a>
  VUID-vkDestroyBufferView-bufferView-parent  
  If `bufferView` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `bufferView` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyBufferView"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
