# VkPipelineCompilerControlCreateInfoAMD(3) Manual Page

## Name

VkPipelineCompilerControlCreateInfoAMD - Structure used to pass compilation control flags to a pipeline



## [](#_c_specification)C Specification

The compilation of a pipeline **can** be tuned by adding a `VkPipelineCompilerControlCreateInfoAMD` structure to the `pNext` chain of [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html) or [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html).

```c++
// Provided by VK_AMD_pipeline_compiler_control
typedef struct VkPipelineCompilerControlCreateInfoAMD {
    VkStructureType                      sType;
    const void*                          pNext;
    VkPipelineCompilerControlFlagsAMD    compilerControlFlags;
} VkPipelineCompilerControlCreateInfoAMD;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `compilerControlFlags` is a bitmask of [VkPipelineCompilerControlFlagBitsAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCompilerControlFlagBitsAMD.html) affecting how the pipeline will be compiled.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPipelineCompilerControlCreateInfoAMD-sType-sType)VUID-VkPipelineCompilerControlCreateInfoAMD-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_COMPILER_CONTROL_CREATE_INFO_AMD`
- [](#VUID-VkPipelineCompilerControlCreateInfoAMD-compilerControlFlags-zerobitmask)VUID-VkPipelineCompilerControlCreateInfoAMD-compilerControlFlags-zerobitmask  
  `compilerControlFlags` **must** be `0`

## [](#_see_also)See Also

[VK\_AMD\_pipeline\_compiler\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_pipeline_compiler_control.html), [VkPipelineCompilerControlFlagsAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCompilerControlFlagsAMD.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineCompilerControlCreateInfoAMD).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0