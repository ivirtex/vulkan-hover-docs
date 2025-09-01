# VkShaderModule(3) Manual Page

## Name

VkShaderModule - Opaque handle to a shader module object



## [](#_c_specification)C Specification

*Shader modules* contain *shader code* and one or more entry points. Shaders are selected from a shader module by specifying an entry point as part of [pipeline](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines) creation. The stages of a pipeline **can** use shaders that come from different modules. The shader code defining a shader module **must** be in the SPIR-V format, as described by the [Vulkan Environment for SPIR-V](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv) appendix.

Shader modules are represented by `VkShaderModule` handles:

```c++
// Provided by VK_VERSION_1_0
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkShaderModule)
```

## [](#_description)Description

## [](#_see_also)See Also

[VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDataGraphPipelineShaderModuleCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineShaderModuleCreateInfoARM.html), [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html), [vkCreateShaderModule](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateShaderModule.html), [vkDestroyShaderModule](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyShaderModule.html), [vkGetShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetShaderModuleIdentifierEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkShaderModule).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0