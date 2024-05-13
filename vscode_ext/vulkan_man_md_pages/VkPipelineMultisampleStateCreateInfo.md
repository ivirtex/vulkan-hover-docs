# VkPipelineMultisampleStateCreateInfo(3) Manual Page

## Name

VkPipelineMultisampleStateCreateInfo - Structure specifying parameters
of a newly created pipeline multisample state



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineMultisampleStateCreateInfo` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `rasterizationSamples` is a
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value specifying
  the number of samples used in rasterization. This value is ignored for
  the purposes of setting the number of samples used in rasterization if
  the pipeline is created with the
  `VK_DYNAMIC_STATE_RASTERIZATION_SAMPLES_EXT` dynamic state set, but if
  `VK_DYNAMIC_STATE_SAMPLE_MASK_EXT` dynamic state is not set, it is
  still used to define the size of the `pSampleMask` array as described
  below.

- `sampleShadingEnable` **can** be used to enable <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-sampleshading"
  target="_blank" rel="noopener">Sample Shading</a>.

- `minSampleShading` specifies a minimum fraction of sample shading if
  `sampleShadingEnable` is set to `VK_TRUE`.

- `pSampleMask` is a pointer to an array of
  [VkSampleMask](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleMask.html) values used in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-samplemask"
  target="_blank" rel="noopener">sample mask test</a>.

- `alphaToCoverageEnable` controls whether a temporary coverage value is
  generated based on the alpha component of the fragment’s first color
  output as specified in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-covg"
  target="_blank" rel="noopener">Multisample Coverage</a> section.

- `alphaToOneEnable` controls whether the alpha component of the
  fragment’s first color output is replaced with one as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-covg"
  target="_blank" rel="noopener">Multisample Coverage</a>.

## <a href="#_description" class="anchor"></a>Description

Each bit in the sample mask is associated with a unique <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-multisampling-coverage-mask"
target="_blank" rel="noopener">sample index</a> as defined for the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-multisampling-coverage-mask"
target="_blank" rel="noopener">coverage mask</a>. Each bit b for mask
word w in the sample mask corresponds to sample index i, where i = 32 ×
w + b. `pSampleMask` has a length equal to ⌈ `rasterizationSamples` / 32
⌉ words.

If `pSampleMask` is `NULL`, it is treated as if the mask has all bits
set to `1`.

Valid Usage

- <a
  href="#VUID-VkPipelineMultisampleStateCreateInfo-sampleShadingEnable-00784"
  id="VUID-VkPipelineMultisampleStateCreateInfo-sampleShadingEnable-00784"></a>
  VUID-VkPipelineMultisampleStateCreateInfo-sampleShadingEnable-00784  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-sampleRateShading"
  target="_blank" rel="noopener"><code>sampleRateShading</code></a>
  feature is not enabled, `sampleShadingEnable` **must** be `VK_FALSE`

- <a
  href="#VUID-VkPipelineMultisampleStateCreateInfo-alphaToOneEnable-00785"
  id="VUID-VkPipelineMultisampleStateCreateInfo-alphaToOneEnable-00785"></a>
  VUID-VkPipelineMultisampleStateCreateInfo-alphaToOneEnable-00785  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-alphaToOne"
  target="_blank" rel="noopener"><code>alphaToOne</code></a> feature is
  not enabled, `alphaToOneEnable` **must** be `VK_FALSE`

- <a
  href="#VUID-VkPipelineMultisampleStateCreateInfo-minSampleShading-00786"
  id="VUID-VkPipelineMultisampleStateCreateInfo-minSampleShading-00786"></a>
  VUID-VkPipelineMultisampleStateCreateInfo-minSampleShading-00786  
  `minSampleShading` **must** be in the range \[0,1\]

- <a
  href="#VUID-VkPipelineMultisampleStateCreateInfo-rasterizationSamples-01415"
  id="VUID-VkPipelineMultisampleStateCreateInfo-rasterizationSamples-01415"></a>
  VUID-VkPipelineMultisampleStateCreateInfo-rasterizationSamples-01415  
  If the
  [`VK_NV_framebuffer_mixed_samples`](VK_NV_framebuffer_mixed_samples.html)
  extension is enabled, and if the subpass has any color attachments and
  `rasterizationSamples` is greater than the number of color samples,
  then `sampleShadingEnable` **must** be `VK_FALSE`

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineMultisampleStateCreateInfo-sType-sType"
  id="VUID-VkPipelineMultisampleStateCreateInfo-sType-sType"></a>
  VUID-VkPipelineMultisampleStateCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_MULTISAMPLE_STATE_CREATE_INFO`

- <a href="#VUID-VkPipelineMultisampleStateCreateInfo-pNext-pNext"
  id="VUID-VkPipelineMultisampleStateCreateInfo-pNext-pNext"></a>
  VUID-VkPipelineMultisampleStateCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkPipelineCoverageModulationStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageModulationStateCreateInfoNV.html),
  [VkPipelineCoverageReductionStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageReductionStateCreateInfoNV.html),
  [VkPipelineCoverageToColorStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageToColorStateCreateInfoNV.html),
  or
  [VkPipelineSampleLocationsStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html)

- <a href="#VUID-VkPipelineMultisampleStateCreateInfo-sType-unique"
  id="VUID-VkPipelineMultisampleStateCreateInfo-sType-unique"></a>
  VUID-VkPipelineMultisampleStateCreateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkPipelineMultisampleStateCreateInfo-flags-zerobitmask"
  id="VUID-VkPipelineMultisampleStateCreateInfo-flags-zerobitmask"></a>
  VUID-VkPipelineMultisampleStateCreateInfo-flags-zerobitmask  
  `flags` **must** be `0`

- <a
  href="#VUID-VkPipelineMultisampleStateCreateInfo-rasterizationSamples-parameter"
  id="VUID-VkPipelineMultisampleStateCreateInfo-rasterizationSamples-parameter"></a>
  VUID-VkPipelineMultisampleStateCreateInfo-rasterizationSamples-parameter  
  `rasterizationSamples` **must** be a valid
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value

- <a
  href="#VUID-VkPipelineMultisampleStateCreateInfo-pSampleMask-parameter"
  id="VUID-VkPipelineMultisampleStateCreateInfo-pSampleMask-parameter"></a>
  VUID-VkPipelineMultisampleStateCreateInfo-pSampleMask-parameter  
  If `pSampleMask` is not `NULL`, `pSampleMask` **must** be a valid
  pointer to an array of ⌈32rasterizationSamples​⌉
  [VkSampleMask](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleMask.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html),
[VkPipelineMultisampleStateCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateFlags.html),
[VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html),
[VkSampleMask](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleMask.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineMultisampleStateCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
