# vkDestroyShaderModule(3) Manual Page

## Name

vkDestroyShaderModule - Destroy a shader module



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a shader module, call:

``` c
// Provided by VK_VERSION_1_0
void vkDestroyShaderModule(
    VkDevice                                    device,
    VkShaderModule                              shaderModule,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the shader module.

- `shaderModule` is the handle of the shader module to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

A shader module **can** be destroyed while pipelines created using its
shaders are still in use.

Valid Usage

- <a href="#VUID-vkDestroyShaderModule-shaderModule-01092"
  id="VUID-vkDestroyShaderModule-shaderModule-01092"></a>
  VUID-vkDestroyShaderModule-shaderModule-01092  
  If `VkAllocationCallbacks` were provided when `shaderModule` was
  created, a compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroyShaderModule-shaderModule-01093"
  id="VUID-vkDestroyShaderModule-shaderModule-01093"></a>
  VUID-vkDestroyShaderModule-shaderModule-01093  
  If no `VkAllocationCallbacks` were provided when `shaderModule` was
  created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyShaderModule-device-parameter"
  id="VUID-vkDestroyShaderModule-device-parameter"></a>
  VUID-vkDestroyShaderModule-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyShaderModule-shaderModule-parameter"
  id="VUID-vkDestroyShaderModule-shaderModule-parameter"></a>
  VUID-vkDestroyShaderModule-shaderModule-parameter  
  If `shaderModule` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `shaderModule` **must** be a valid
  [VkShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModule.html) handle

- <a href="#VUID-vkDestroyShaderModule-pAllocator-parameter"
  id="VUID-vkDestroyShaderModule-pAllocator-parameter"></a>
  VUID-vkDestroyShaderModule-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyShaderModule-shaderModule-parent"
  id="VUID-vkDestroyShaderModule-shaderModule-parent"></a>
  VUID-vkDestroyShaderModule-shaderModule-parent  
  If `shaderModule` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `shaderModule` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModule.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyShaderModule"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
