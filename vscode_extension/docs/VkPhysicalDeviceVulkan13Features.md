# VkPhysicalDeviceVulkan13Features(3) Manual Page

## Name

VkPhysicalDeviceVulkan13Features - Structure describing the Vulkan 1.3
features that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceVulkan13Features` structure is defined as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkPhysicalDeviceVulkan13Features {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           robustImageAccess;
    VkBool32           inlineUniformBlock;
    VkBool32           descriptorBindingInlineUniformBlockUpdateAfterBind;
    VkBool32           pipelineCreationCacheControl;
    VkBool32           privateData;
    VkBool32           shaderDemoteToHelperInvocation;
    VkBool32           shaderTerminateInvocation;
    VkBool32           subgroupSizeControl;
    VkBool32           computeFullSubgroups;
    VkBool32           synchronization2;
    VkBool32           textureCompressionASTC_HDR;
    VkBool32           shaderZeroInitializeWorkgroupMemory;
    VkBool32           dynamicRendering;
    VkBool32           shaderIntegerDotProduct;
    VkBool32           maintenance4;
} VkPhysicalDeviceVulkan13Features;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="features-robustImageAccess"></span> `robustImageAccess`
  indicates whether image accesses are tightly bounds-checked against
  the dimensions of the image view. [Invalid
  texels](#textures-input-validation) resulting from out of bounds image
  loads will be replaced as described in [Texel
  Replacement](#textures-texel-replacement), with either (0,0,1) or
  (0,0,0) values inserted for missing G, B, or A components based on the
  format.

- <span id="features-inlineUniformBlock"></span> `inlineUniformBlock`
  indicates whether the implementation supports inline uniform block
  descriptors. If this feature is not enabled,
  `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` **must** not be used.

- <span id="features-descriptorBindingInlineUniformBlockUpdateAfterBind"></span>
  `descriptorBindingInlineUniformBlockUpdateAfterBind` indicates whether
  the implementation supports updating inline uniform block descriptors
  after a set is bound. If this feature is not enabled,
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` **must** not be used
  with `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`.

- <span id="features-pipelineCreationCacheControl"></span>
  `pipelineCreationCacheControl` indicates that the implementation
  supports:

  - The following **can** be used in `Vk*PipelineCreateInfo`::`flags`:

    - `VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT`

    - `VK_PIPELINE_CREATE_EARLY_RETURN_ON_FAILURE_BIT`

  - The following **can** be used in
    [VkPipelineCacheCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCacheCreateInfo.html)::`flags`:

    - `VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT`

- <span id="features-privateData"></span> `privateData` indicates
  whether the implementation supports private data. See [Private
  Data](#private-data).

- <span id="features-shaderDemoteToHelperInvocation"></span>
  `shaderDemoteToHelperInvocation` indicates whether the implementation
  supports the SPIR-V `DemoteToHelperInvocationEXT` capability.

- <span id="features-shaderTerminateInvocation"></span>
  `shaderTerminateInvocation` specifies whether the implementation
  supports SPIR-V modules that use the `SPV_KHR_terminate_invocation`
  extension.

- <span id="features-subgroupSizeControl"></span> `subgroupSizeControl`
  indicates whether the implementation supports controlling shader
  subgroup sizes via the
  `VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT` flag
  and the
  [VkPipelineShaderStageRequiredSubgroupSizeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfo.html)
  structure.

- <span id="features-computeFullSubgroups"></span>
  `computeFullSubgroups` indicates whether the implementation supports
  requiring full subgroups in compute , mesh, or task shaders via the
  `VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT` flag.

- <span id="features-synchronization2"></span> `synchronization2`
  indicates whether the implementation supports the new set of
  synchronization commands introduced in
  [`VK_KHR_synchronization2`](VK_KHR_synchronization2.html).

- <span id="features-textureCompressionASTC_HDR"></span>
  `textureCompressionASTC_HDR` indicates whether all of the ASTC HDR
  compressed texture formats are supported. If this feature is enabled,
  then the `VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT`,
  `VK_FORMAT_FEATURE_BLIT_SRC_BIT` and
  `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT` features **must**
  be supported in `optimalTilingFeatures` for the following formats:

  - `VK_FORMAT_ASTC_4x4_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_5x4_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_5x5_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_6x5_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_6x6_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_8x5_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_8x6_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_8x8_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_10x5_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_10x6_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_10x8_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_10x10_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_12x10_SFLOAT_BLOCK`

  - `VK_FORMAT_ASTC_12x12_SFLOAT_BLOCK`

    To query for additional properties, or if the feature is not
    enabled,
    [vkGetPhysicalDeviceFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties.html)
    and
    [vkGetPhysicalDeviceImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties.html)
    **can** be used to check for supported properties of individual
    formats as normal.

- <span id="features-shaderZeroInitializeWorkgroupMemory"></span>
  `shaderZeroInitializeWorkgroupMemory` specifies whether the
  implementation supports initializing a variable in Workgroup storage
  class.

- <span id="features-dynamicRendering"></span> `dynamicRendering`
  specifies that the implementation supports dynamic render pass
  instances using the [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html)
  command.

- <span id="features-shaderIntegerDotProduct"></span>
  `shaderIntegerDotProduct` specifies whether shader modules **can**
  declare the `DotProductInputAllKHR`, `DotProductInput4x8BitKHR`,
  `DotProductInput4x8BitPackedKHR` and `DotProductKHR` capabilities.

- <span id="features-maintenance4"></span> `maintenance4` indicates that
  the implementation supports the following:

  - The application **may** destroy a
    [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) object immediately after
    using it to create another object.

  - `LocalSizeId` **can** be used as an alternative to `LocalSize` to
    specify the local workgroup size with specialization constants.

  - Images created with identical creation parameters will always have
    the same alignment requirements.

  - The size memory requirement of a buffer or image is never greater
    than that of another buffer or image created with a greater or equal
    size.

  - Push constants do not have to be initialized before they are
    dynamically accessed.

  - The interface matching rules allow a larger output vector to match
    with a smaller input vector, with additional values being discarded.

If the `VkPhysicalDeviceVulkan13Features` structure is included in the
`pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceVulkan13Features` **can** also be used in the `pNext`
chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively
enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceVulkan13Features-sType-sType"
  id="VUID-VkPhysicalDeviceVulkan13Features-sType-sType"></a>
  VUID-VkPhysicalDeviceVulkan13Features-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_3_FEATURES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceVulkan13Features"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
