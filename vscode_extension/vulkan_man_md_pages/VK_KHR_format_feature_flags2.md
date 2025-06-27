# VK_KHR_format_feature_flags2(3) Manual Page

## Name

VK_KHR_format_feature_flags2 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

361

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.3-promotions"
  target="_blank" rel="noopener">Vulkan 1.3</a>

## <a href="#_contact" class="anchor"></a>Contact

- Lionel Landwerlin <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_format_feature_flags2%5D%20@llandwerlin%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_format_feature_flags2%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>llandwerlin</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This extension adds a new
[VkFormatFeatureFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits2KHR.html) 64bits
format feature flag type to extend the existing
[VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html) which is limited
to 31 flags. At the time of this writing 29 bits of
[VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html) are already
used.

Because [VkFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties2.html) is already
defined to extend the Vulkan 1.0
[vkGetPhysicalDeviceFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties.html)
command, this extension defines a new
[VkFormatProperties3KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties3KHR.html) to extend the
[VkFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties.html).

On top of replicating all the bits from
[VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html),
[VkFormatFeatureFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits2KHR.html) adds the
following bits :

- `VK_FORMAT_FEATURE_2_STORAGE_READ_WITHOUT_FORMAT_BIT_KHR` and
  `VK_FORMAT_FEATURE_2_STORAGE_WRITE_WITHOUT_FORMAT_BIT_KHR` indicate
  that an implementation supports respectively reading and writing a
  given [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) through storage operations without
  specifying the format in the shader.

- `VK_FORMAT_FEATURE_2_SAMPLED_IMAGE_DEPTH_COMPARISON_BIT_KHR` indicates
  that an implementation supports depth comparison performed by
  `OpImage*Dref*` instructions on a given [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html).
  Previously the result of executing a `OpImage*Dref*` instruction on an
  image view, where the `format` was not one of the depth/stencil
  formats with a depth component, was undefined. This bit clarifies on
  which formats such instructions can be used.

Prior to version 2 of this extension, implementations exposing the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderStorageImageReadWithoutFormat"
target="_blank"
rel="noopener"><code>shaderStorageImageReadWithoutFormat</code></a> and
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderStorageImageWriteWithoutFormat"
target="_blank"
rel="noopener"><code>shaderStorageImageWriteWithoutFormat</code></a>
features may not report
`VK_FORMAT_FEATURE_2_STORAGE_READ_WITHOUT_FORMAT_BIT_KHR` and
`VK_FORMAT_FEATURE_2_STORAGE_WRITE_WITHOUT_FORMAT_BIT_KHR` in
[VkFormatProperties3KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties3KHR.html)::`bufferFeatures`.
Despite this, buffer reads/writes are supported as intended by the
original features.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties2.html):

  - [VkFormatProperties3KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties3KHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkFormatFeatureFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits2KHR.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkFormatFeatureFlags2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlags2KHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_FORMAT_FEATURE_FLAGS_2_EXTENSION_NAME`

- `VK_KHR_FORMAT_FEATURE_FLAGS_2_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_3_KHR`

## <a href="#_promotion_to_vulkan_1_3" class="anchor"></a>Promotion to Vulkan 1.3

Functionality in this extension is included in core Vulkan 1.3, with the
KHR suffix omitted. The original type, enum and command names are still
available as aliases of the core functionality.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 2, 2022-07-20 (Lionel Landwerlin)

  - Clarify that
    VK_FORMAT_FEATURE_2_STORAGE\_(READ\|WRITE)\_WITHOUT_FORMAT_BIT also
    apply to buffer views.

- Revision 1, 2020-07-21 (Lionel Landwerlin)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_format_feature_flags2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
