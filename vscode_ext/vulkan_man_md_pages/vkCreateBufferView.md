# vkCreateBufferView(3) Manual Page

## Name

vkCreateBufferView - Create a new buffer view object



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a buffer view, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkCreateBufferView(
    VkDevice                                    device,
    const VkBufferViewCreateInfo*               pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkBufferView*                               pView);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the buffer view.

- `pCreateInfo` is a pointer to a
  [VkBufferViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferViewCreateInfo.html) structure
  containing parameters to be used to create the buffer view.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pView` is a pointer to a [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) handle in
  which the resulting buffer view object is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkCreateBufferView-device-parameter"
  id="VUID-vkCreateBufferView-device-parameter"></a>
  VUID-vkCreateBufferView-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateBufferView-pCreateInfo-parameter"
  id="VUID-vkCreateBufferView-pCreateInfo-parameter"></a>
  VUID-vkCreateBufferView-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkBufferViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferViewCreateInfo.html) structure

- <a href="#VUID-vkCreateBufferView-pAllocator-parameter"
  id="VUID-vkCreateBufferView-pAllocator-parameter"></a>
  VUID-vkCreateBufferView-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateBufferView-pView-parameter"
  id="VUID-vkCreateBufferView-pView-parameter"></a>
  VUID-vkCreateBufferView-pView-parameter  
  `pView` **must** be a valid pointer to a
  [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html),
[VkBufferViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferViewCreateInfo.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateBufferView"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
