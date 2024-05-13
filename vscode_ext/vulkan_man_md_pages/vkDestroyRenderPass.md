# vkDestroyRenderPass(3) Manual Page

## Name

vkDestroyRenderPass - Destroy a render pass object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a render pass, call:

``` c
// Provided by VK_VERSION_1_0
void vkDestroyRenderPass(
    VkDevice                                    device,
    VkRenderPass                                renderPass,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the render pass.

- `renderPass` is the handle of the render pass to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkDestroyRenderPass-renderPass-00873"
  id="VUID-vkDestroyRenderPass-renderPass-00873"></a>
  VUID-vkDestroyRenderPass-renderPass-00873  
  All submitted commands that refer to `renderPass` **must** have
  completed execution

- <a href="#VUID-vkDestroyRenderPass-renderPass-00874"
  id="VUID-vkDestroyRenderPass-renderPass-00874"></a>
  VUID-vkDestroyRenderPass-renderPass-00874  
  If `VkAllocationCallbacks` were provided when `renderPass` was
  created, a compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroyRenderPass-renderPass-00875"
  id="VUID-vkDestroyRenderPass-renderPass-00875"></a>
  VUID-vkDestroyRenderPass-renderPass-00875  
  If no `VkAllocationCallbacks` were provided when `renderPass` was
  created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyRenderPass-device-parameter"
  id="VUID-vkDestroyRenderPass-device-parameter"></a>
  VUID-vkDestroyRenderPass-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyRenderPass-renderPass-parameter"
  id="VUID-vkDestroyRenderPass-renderPass-parameter"></a>
  VUID-vkDestroyRenderPass-renderPass-parameter  
  If `renderPass` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `renderPass` **must** be a valid [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html)
  handle

- <a href="#VUID-vkDestroyRenderPass-pAllocator-parameter"
  id="VUID-vkDestroyRenderPass-pAllocator-parameter"></a>
  VUID-vkDestroyRenderPass-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyRenderPass-renderPass-parent"
  id="VUID-vkDestroyRenderPass-renderPass-parent"></a>
  VUID-vkDestroyRenderPass-renderPass-parent  
  If `renderPass` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `renderPass` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyRenderPass"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
