# VkPipelineMultisampleStateCreateInfo(3) Manual Page

## Name

VkPipelineMultisampleStateCreateInfo - Structure specifying parameters of a newly created pipeline multisample state



## [](#_c_specification)C Specification

The `VkPipelineMultisampleStateCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkPipelineMultisampleStateCreateInfo {
    VkStructureType                          sType;
    const void*                              pNext;
    VkPipelineMultisampleStateCreateFlags    flags;
    VkSampleCountFlagBits                    rasterizationSamples;
    VkBool32                                 sampleShadingEnable;
    float                                    minSampleShading;
    const VkSampleMask*                      pSampleMask;
    VkBool32                                 alphaToCoverageEnable;
    VkBool32                                 alphaToOneEnable;
} VkPipelineMultisampleStateCreateInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `rasterizationSamples` is a [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) value specifying the number of samples used in rasterization. This value is ignored for the purposes of setting the number of samples used in rasterization if the pipeline is created with the `VK_DYNAMIC_STATE_RASTERIZATION_SAMPLES_EXT` dynamic state set, but if `VK_DYNAMIC_STATE_SAMPLE_MASK_EXT` dynamic state is not set, it is still used to define the size of the `pSampleMask` array as described below.
- `sampleShadingEnable` **can** be used to enable [Sample Shading](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-sampleshading).
- `minSampleShading` specifies a minimum fraction of sample shading if `sampleShadingEnable` is `VK_TRUE`.
- `pSampleMask` is a pointer to an array of [VkSampleMask](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleMask.html) values used in the [sample mask test](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-samplemask).
- `alphaToCoverageEnable` controls whether a temporary coverage value is generated based on the alpha component of the fragment’s first color output as specified in the [Multisample Coverage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-covg) section.
- `alphaToOneEnable` controls whether the alpha component of the fragment’s first color output is replaced with one as described in [Multisample Coverage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-covg).

## [](#_description)Description

Each bit in the sample mask is associated with a unique [sample index](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-multisampling-coverage-mask) as defined for the [coverage mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-multisampling-coverage-mask). Each bit b for mask word w in the sample mask corresponds to sample index i, where i = 32 × w + b. `pSampleMask` has a length equal to ⌈ `rasterizationSamples` / 32 ⌉ words.

If `pSampleMask` is `NULL`, it is treated as if the mask has all bits set to `1`.

Valid Usage

- [](#VUID-VkPipelineMultisampleStateCreateInfo-sampleShadingEnable-00784)VUID-VkPipelineMultisampleStateCreateInfo-sampleShadingEnable-00784  
  If the [`sampleRateShading`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-sampleRateShading) feature is not enabled, `sampleShadingEnable` **must** be `VK_FALSE`
- [](#VUID-VkPipelineMultisampleStateCreateInfo-alphaToOneEnable-00785)VUID-VkPipelineMultisampleStateCreateInfo-alphaToOneEnable-00785  
  If the [`alphaToOne`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-alphaToOne) feature is not enabled, `alphaToOneEnable` **must** be `VK_FALSE`
- [](#VUID-VkPipelineMultisampleStateCreateInfo-minSampleShading-00786)VUID-VkPipelineMultisampleStateCreateInfo-minSampleShading-00786  
  `minSampleShading` **must** be in the range \[0,1]
- [](#VUID-VkPipelineMultisampleStateCreateInfo-rasterizationSamples-01415)VUID-VkPipelineMultisampleStateCreateInfo-rasterizationSamples-01415  
  If the `VK_NV_framebuffer_mixed_samples` extension is enabled, and the [`coverageReductionMode`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-coverageReductionMode) feature is not enabled, or the `pNext` chain does not contain `VkPipelineCoverageReductionStateCreateInfoNV`, or `VkPipelineCoverageReductionStateCreateInfoNV`::`coverageReductionMode` is not set to `VK_COVERAGE_REDUCTION_MODE_TRUNCATE_NV`, and the subpass has any color attachments, and `rasterizationSamples` is greater than the number of color samples, then [sample shading](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-sampleshading) **must** not be enabled

Valid Usage (Implicit)

- [](#VUID-VkPipelineMultisampleStateCreateInfo-sType-sType)VUID-VkPipelineMultisampleStateCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_MULTISAMPLE_STATE_CREATE_INFO`
- [](#VUID-VkPipelineMultisampleStateCreateInfo-pNext-pNext)VUID-VkPipelineMultisampleStateCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkPipelineCoverageModulationStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCoverageModulationStateCreateInfoNV.html), [VkPipelineCoverageReductionStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCoverageReductionStateCreateInfoNV.html), [VkPipelineCoverageToColorStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCoverageToColorStateCreateInfoNV.html), or [VkPipelineSampleLocationsStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html)
- [](#VUID-VkPipelineMultisampleStateCreateInfo-sType-unique)VUID-VkPipelineMultisampleStateCreateInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkPipelineMultisampleStateCreateInfo-flags-zerobitmask)VUID-VkPipelineMultisampleStateCreateInfo-flags-zerobitmask  
  `flags` **must** be `0`
- [](#VUID-VkPipelineMultisampleStateCreateInfo-rasterizationSamples-parameter)VUID-VkPipelineMultisampleStateCreateInfo-rasterizationSamples-parameter  
  `rasterizationSamples` **must** be a valid [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) value

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkPipelineMultisampleStateCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineMultisampleStateCreateFlags.html), [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html), [VkSampleMask](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleMask.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineMultisampleStateCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0