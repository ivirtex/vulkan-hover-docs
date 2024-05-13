# VkPipelineCompilerControlCreateInfoAMD(3) Manual Page

## Name

VkPipelineCompilerControlCreateInfoAMD - Structure used to pass
compilation control flags to a pipeline



## <a href="#_c_specification" class="anchor"></a>C Specification

The compilation of a pipeline **can** be tuned by adding a
`VkPipelineCompilerControlCreateInfoAMD` structure to the `pNext` chain
of [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html) or
[VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineCreateInfo.html).

``` c
// Provided by VK_AMD_pipeline_compiler_control
typedef struct VkPipelineCompilerControlCreateInfoAMD {
    VkStructureType                      sType;
    const void*                          pNext;
    VkPipelineCompilerControlFlagsAMD    compilerControlFlags;
} VkPipelineCompilerControlCreateInfoAMD;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `compilerControlFlags` is a bitmask of
  [VkPipelineCompilerControlFlagBitsAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCompilerControlFlagBitsAMD.html)
  affecting how the pipeline will be compiled.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineCompilerControlCreateInfoAMD-sType-sType"
  id="VUID-VkPipelineCompilerControlCreateInfoAMD-sType-sType"></a>
  VUID-VkPipelineCompilerControlCreateInfoAMD-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_COMPILER_CONTROL_CREATE_INFO_AMD`

- <a
  href="#VUID-VkPipelineCompilerControlCreateInfoAMD-compilerControlFlags-zerobitmask"
  id="VUID-VkPipelineCompilerControlCreateInfoAMD-compilerControlFlags-zerobitmask"></a>
  VUID-VkPipelineCompilerControlCreateInfoAMD-compilerControlFlags-zerobitmask  
  `compilerControlFlags` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_AMD_pipeline_compiler_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMD_pipeline_compiler_control.html),
[VkPipelineCompilerControlFlagsAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCompilerControlFlagsAMD.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineCompilerControlCreateInfoAMD"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
