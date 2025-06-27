# VkShaderModule(3) Manual Page

## Name

VkShaderModule - Opaque handle to a shader module object



## <a href="#_c_specification" class="anchor"></a>C Specification

*Shader modules* contain *shader code* and one or more entry points.
Shaders are selected from a shader module by specifying an entry point
as part of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines"
target="_blank" rel="noopener">pipeline</a> creation. The stages of a
pipeline **can** use shaders that come from different modules. The
shader code defining a shader module **must** be in the SPIR-V format,
as described by the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv"
target="_blank" rel="noopener">Vulkan Environment for SPIR-V</a>
appendix.

Shader modules are represented by `VkShaderModule` handles:

``` c
// Provided by VK_VERSION_1_0
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkShaderModule)
```

## <a href="#_description" class="anchor"></a>Description

## <a href="#_see_also" class="anchor"></a>See Also

[VK_DEFINE_NON_DISPATCHABLE_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html),
[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html),
[vkCreateShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateShaderModule.html),
[vkDestroyShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyShaderModule.html),
[vkGetShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetShaderModuleIdentifierEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkShaderModule"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
