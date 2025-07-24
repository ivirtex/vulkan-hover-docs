# vkCreateShaderModule(3) Manual Page

## Name

vkCreateShaderModule - Creates a new shader module object



## [](#_c_specification)C Specification

To create a shader module, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkCreateShaderModule(
    VkDevice                                    device,
    const VkShaderModuleCreateInfo*             pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkShaderModule*                             pShaderModule);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the shader module.
- `pCreateInfo` is a pointer to a [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html) structure.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pShaderModule` is a pointer to a [VkShaderModule](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModule.html) handle in which the resulting shader module object is returned.

## [](#_description)Description

Once a shader module has been created, any entry points it contains **can** be used in pipeline shader stages as described in [Compute Pipelines](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-compute) and [Graphics Pipelines](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-graphics).

Note

If the [`maintenance5`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance5) feature is enabled, shader module creation can be omitted entirely. Instead, applications should provide the [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html) structure directly in to pipeline creation by chaining it to [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html). This avoids the overhead of creating and managing an additional object.

Valid Usage

- [](#VUID-vkCreateShaderModule-pCreateInfo-06904)VUID-vkCreateShaderModule-pCreateInfo-06904  
  If `pCreateInfo` is not `NULL`, `pCreateInfo->pNext` **must** be `NULL` or a pointer to a valid instance of
  
  - [VkShaderModuleValidationCacheCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleValidationCacheCreateInfoEXT.html)
  - [VkValidationFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFeaturesEXT.html)

Valid Usage (Implicit)

- [](#VUID-vkCreateShaderModule-device-parameter)VUID-vkCreateShaderModule-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateShaderModule-pCreateInfo-parameter)VUID-vkCreateShaderModule-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html) structure
- [](#VUID-vkCreateShaderModule-pAllocator-parameter)VUID-vkCreateShaderModule-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateShaderModule-pShaderModule-parameter)VUID-vkCreateShaderModule-pShaderModule-parameter  
  `pShaderModule` **must** be a valid pointer to a [VkShaderModule](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModule.html) handle

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INVALID_SHADER_NV`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkShaderModule](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModule.html), [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateShaderModule)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0