# vkDestroyShaderEXT(3) Manual Page

## Name

vkDestroyShaderEXT - Destroy a shader object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a shader object, call:

``` c
// Provided by VK_EXT_shader_object
void vkDestroyShaderEXT(
    VkDevice                                    device,
    VkShaderEXT                                 shader,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the shader object.

- `shader` is the handle of the shader object to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Destroying a shader object used by one or more command buffers in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
target="_blank" rel="noopener">recording or executable state</a> causes
those command buffers to move into the *invalid state*.

Valid Usage

- <a href="#VUID-vkDestroyShaderEXT-None-08481"
  id="VUID-vkDestroyShaderEXT-None-08481"></a>
  VUID-vkDestroyShaderEXT-None-08481  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderObject"
  target="_blank" rel="noopener"><code>shaderObject</code></a> feature
  **must** be enabled

- <a href="#VUID-vkDestroyShaderEXT-shader-08482"
  id="VUID-vkDestroyShaderEXT-shader-08482"></a>
  VUID-vkDestroyShaderEXT-shader-08482  
  All submitted commands that refer to `shader` **must** have completed
  execution

- <a href="#VUID-vkDestroyShaderEXT-pAllocator-08483"
  id="VUID-vkDestroyShaderEXT-pAllocator-08483"></a>
  VUID-vkDestroyShaderEXT-pAllocator-08483  
  If `VkAllocationCallbacks` were provided when `shader` was created, a
  compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroyShaderEXT-pAllocator-08484"
  id="VUID-vkDestroyShaderEXT-pAllocator-08484"></a>
  VUID-vkDestroyShaderEXT-pAllocator-08484  
  If no `VkAllocationCallbacks` were provided when `shader` was created,
  `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyShaderEXT-device-parameter"
  id="VUID-vkDestroyShaderEXT-device-parameter"></a>
  VUID-vkDestroyShaderEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyShaderEXT-shader-parameter"
  id="VUID-vkDestroyShaderEXT-shader-parameter"></a>
  VUID-vkDestroyShaderEXT-shader-parameter  
  If `shader` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `shader`
  **must** be a valid [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html) handle

- <a href="#VUID-vkDestroyShaderEXT-pAllocator-parameter"
  id="VUID-vkDestroyShaderEXT-pAllocator-parameter"></a>
  VUID-vkDestroyShaderEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyShaderEXT-shader-parent"
  id="VUID-vkDestroyShaderEXT-shader-parent"></a>
  VUID-vkDestroyShaderEXT-shader-parent  
  If `shader` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `shader` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_shader_object](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_object.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyShaderEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
