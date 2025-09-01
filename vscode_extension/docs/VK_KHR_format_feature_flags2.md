# VK\_KHR\_format\_feature\_flags2(3) Manual Page

## Name

VK\_KHR\_format\_feature\_flags2 - device extension



## [](#_registered_extension_number)Registered Extension Number

361

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_2
- Interacts with VK\_EXT\_filter\_cubic
- Interacts with VK\_EXT\_sampler\_filter\_minmax
- Interacts with VK\_IMG\_filter\_cubic

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.3](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.3-promotions)

## [](#_contact)Contact

- Lionel Landwerlin [\[GitHub\]llandwerlin](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_format_feature_flags2%5D%20%40llandwerlin%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_format_feature_flags2%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-07-01

**IP Status**

No known IP claims.

**Contributors**

- Lionel Landwerlin, Intel
- Faith Ekstrand, Intel
- Tobias Hector, AMD
- Spencer Fricke, Samsung Electronics
- Graeme Leese, Broadcom
- Jan-Harald Fredriksen, ARM

## [](#_description)Description

This extension adds a new [VkFormatFeatureFlagBits2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2KHR.html) 64bits format feature flag type to extend the existing [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits.html) which is limited to 31 flags. At the time of this writing 29 bits of [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits.html) are already used.

Because [VkFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties2.html) is already defined to extend the Vulkan 1.0 [vkGetPhysicalDeviceFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFormatProperties.html) command, this extension defines a new [VkFormatProperties3KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties3KHR.html) to extend the [VkFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties.html).

On top of replicating all the bits from [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits.html), [VkFormatFeatureFlagBits2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2KHR.html) adds the following bits :

- `VK_FORMAT_FEATURE_2_STORAGE_READ_WITHOUT_FORMAT_BIT_KHR` and `VK_FORMAT_FEATURE_2_STORAGE_WRITE_WITHOUT_FORMAT_BIT_KHR` specify that an implementation supports reading and writing, respectively, a given [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) through storage operations without specifying the format in the shader.
- `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_DEPTH_COMPARISON_BIT_KHR` specifies that an implementation supports depth comparison performed by `OpImage*Dref*` instructions on a given [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html). Previously the result of executing a `OpImage*Dref*` instruction on an image view, where the `format` was not one of the depth/stencil formats with a depth component, was undefined. This bit clarifies on which formats such instructions can be used.

Prior to version 2 of this extension, implementations exposing the [`shaderStorageImageReadWithoutFormat`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderStorageImageReadWithoutFormat) and [`shaderStorageImageWriteWithoutFormat`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderStorageImageWriteWithoutFormat) features may not report `VK_FORMAT_FEATURE_2_STORAGE_READ_WITHOUT_FORMAT_BIT_KHR` and `VK_FORMAT_FEATURE_2_STORAGE_WRITE_WITHOUT_FORMAT_BIT_KHR` in [VkFormatProperties3KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties3KHR.html)::`bufferFeatures`. Despite this, buffer reads/writes are supported as intended by the original features.

## [](#_new_structures)New Structures

- Extending [VkFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties2.html):
  
  - [VkFormatProperties3KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties3KHR.html)

## [](#_new_enums)New Enums

- [VkFormatFeatureFlagBits2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2KHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkFormatFeatureFlags2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlags2KHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_FORMAT_FEATURE_FLAGS_2_EXTENSION_NAME`
- `VK_KHR_FORMAT_FEATURE_FLAGS_2_SPEC_VERSION`
- Extending [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2.html):
  
  - `VK_FORMAT_FEATURE_2_BLIT_DST_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_BLIT_SRC_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_COLOR_ATTACHMENT_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_COLOR_ATTACHMENT_BLEND_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_COSITED_CHROMA_SAMPLES_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_DEPTH_STENCIL_ATTACHMENT_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_DISJOINT_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_MIDPOINT_CHROMA_SAMPLES_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_DEPTH_COMPARISON_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_FILTER_LINEAR_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_STORAGE_IMAGE_ATOMIC_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_STORAGE_IMAGE_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_STORAGE_READ_WITHOUT_FORMAT_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_STORAGE_TEXEL_BUFFER_ATOMIC_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_STORAGE_TEXEL_BUFFER_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_STORAGE_WRITE_WITHOUT_FORMAT_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_TRANSFER_DST_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_TRANSFER_SRC_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_UNIFORM_TEXEL_BUFFER_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_VERTEX_BUFFER_BIT_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_3_KHR`

If [VK\_EXT\_filter\_cubic](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_filter_cubic.html) or [VK\_IMG\_filter\_cubic](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_IMG_filter_cubic.html) is supported:

- Extending [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2.html):
  
  - `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_FILTER_CUBIC_BIT_EXT`

If [Vulkan Version 1.2](#versions-1.2) or [VK\_EXT\_sampler\_filter\_minmax](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_sampler_filter_minmax.html) is supported:

- Extending [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2.html):
  
  - `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_FILTER_MINMAX_BIT_KHR`

## [](#_promotion_to_vulkan_1_3)Promotion to Vulkan 1.3

Vulkan APIs in this extension are included in core Vulkan 1.3, with the KHR suffix omitted. External interactions defined by this extension, such as SPIR-V token names, retain their original names. The original Vulkan API names are still available as aliases of the core functionality.

## [](#_version_history)Version History

- Revision 2, 2022-07-20 (Lionel Landwerlin)
  
  - Clarify that VK\_FORMAT\_FEATURE\_2\_STORAGE\_(READ|WRITE)\_WITHOUT\_FORMAT\_BIT also apply to buffer views.
- Revision 1, 2020-07-21 (Lionel Landwerlin)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_format_feature_flags2).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0