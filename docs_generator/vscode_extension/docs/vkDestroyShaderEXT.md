# vkDestroyShaderEXT(3) Manual Page

## Name

vkDestroyShaderEXT - Destroy a shader object



## [](#_c_specification)C Specification

To destroy a shader object, call:

```c++
// Provided by VK_EXT_shader_object
void vkDestroyShaderEXT(
    VkDevice                                    device,
    VkShaderEXT                                 shader,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the shader object.
- `shader` is the handle of the shader object to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Destroying a shader object used by one or more command buffers in the [recording or executable state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle) causes those command buffers to move into the *invalid state*.

Valid Usage

- [](#VUID-vkDestroyShaderEXT-None-08481)VUID-vkDestroyShaderEXT-None-08481  
  The [`shaderObject`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderObject) feature **must** be enabled
- [](#VUID-vkDestroyShaderEXT-shader-08482)VUID-vkDestroyShaderEXT-shader-08482  
  All submitted commands that refer to `shader` **must** have completed execution
- [](#VUID-vkDestroyShaderEXT-pAllocator-08483)VUID-vkDestroyShaderEXT-pAllocator-08483  
  If `VkAllocationCallbacks` were provided when `shader` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyShaderEXT-pAllocator-08484)VUID-vkDestroyShaderEXT-pAllocator-08484  
  If no `VkAllocationCallbacks` were provided when `shader` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroyShaderEXT-device-parameter)VUID-vkDestroyShaderEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyShaderEXT-shader-parameter)VUID-vkDestroyShaderEXT-shader-parameter  
  If `shader` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `shader` **must** be a valid [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) handle
- [](#VUID-vkDestroyShaderEXT-pAllocator-parameter)VUID-vkDestroyShaderEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyShaderEXT-shader-parent)VUID-vkDestroyShaderEXT-shader-parent  
  If `shader` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `shader` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyShaderEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0