# VK_NV_fragment_shading_rate_enums(3) Manual Page

## Name

VK_NV_fragment_shading_rate_enums - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

327

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_fragment_shading_rate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_fragment_shading_rate.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Pat Brown <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_fragment_shading_rate_enums%5D%20@nvpbrown%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_fragment_shading_rate_enums%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>nvpbrown</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-09-02

**Contributors**  
- Pat Brown, NVIDIA

- Jeff Bolz, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension builds on the fragment shading rate functionality
provided by the VK_KHR_fragment_shading_rate extension, adding support
for “supersample” fragment shading rates that trigger multiple fragment
shader invocations per pixel as well as a “no invocations” shading rate
that discards any portions of a primitive that would use that shading
rate.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdSetFragmentShadingRateEnumNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetFragmentShadingRateEnumNV.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html):

  - [VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceFragmentShadingRateEnumsFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentShadingRateEnumsFeaturesNV.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkFragmentShadingRateNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateNV.html)

- [VkFragmentShadingRateTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateTypeNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_FRAGMENT_SHADING_RATE_ENUMS_EXTENSION_NAME`

- `VK_NV_FRAGMENT_SHADING_RATE_ENUMS_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADING_RATE_ENUMS_FEATURES_NV`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADING_RATE_ENUMS_PROPERTIES_NV`

  - `VK_STRUCTURE_TYPE_PIPELINE_FRAGMENT_SHADING_RATE_ENUM_STATE_CREATE_INFO_NV`

## <a href="#_issues" class="anchor"></a>Issues

1.  Why was this extension created? How should it be named?

    **RESOLVED**: The primary goal of this extension was to expose
    support for supersample and “no invocations” shading rates, which
    are supported by the VK_NV_shading_rate_image extension but not by
    VK_KHR_fragment_shading_rate. Because VK_KHR_fragment_shading_rate
    specifies the primitive shading rate using a fragment size in
    pixels, it lacks a good way to specify supersample rates. To deal
    with this, we defined enums covering shading rates supported by the
    KHR extension as well as the new shading rates and added structures
    and APIs accepting shading rate enums instead of fragment sizes.

    Since this extension adds two different types of shading rates, both
    expressed using enums, we chose the extension name
    VK_NV_fragment_shading_rate_enums.

2.  Is this a standalone extension?

    **RESOLVED**: No, this extension requires
    VK_KHR_fragment_shading_rate. In order to use the features of this
    extension, applications must enable the relevant features of KHR
    extension.

3.  How are the shading rate enums used, and how were the enum values
    assigned?

    **RESOLVED**: The shading rates supported by the enums in this
    extension are accepted as pipeline, primitive, and attachment
    shading rates and behave identically. For the shading rates also
    supported by the KHR extension, the values assigned to the
    corresponding enums are identical to the values already used for the
    primitive and attachment shading rates in the KHR extension. For
    those enums, bits 0 and 1 specify the base two logarithm of the
    fragment height and bits 2 and 3 specify the base two logarithm of
    the fragment width. For the new shading rates added by this
    extension, we chose to use 11 through 14 (10 plus the base two
    logarithm of the invocation count) for the supersample rates and 15
    for the “no invocations” rate. None of those values are supported as
    primitive or attachment shading rates by the KHR extension.

4.  Between this extension, VK_KHR_fragment_shading_rate, and
    VK_NV_shading_rate_image, there are three different ways to specify
    shading rate state in a pipeline. How should we handle this?

    **RESOLVED**: We do not allow the concurrent use of
    VK_NV_shading_rate_image and VK_KHR_fragment_shading_rate; it is an
    error to enable shading rate features from both extensions. But we
    do allow applications to enable this extension together with
    VK_KHR_fragment_shading_rate together. While we expect that
    applications will never attach pipeline CreateInfo structures for
    both this extension and the KHR extension concurrently, Vulkan does
    not have any precedent forbidding such behavior and instead
    typically treats a pipeline created without an extension-specific
    CreateInfo structure as equivalent to one containing default values
    specified by the extension. Rather than adding such a rule
    considering the presence or absence of our new CreateInfo structure,
    we instead included a `shadingRateType` member to
    [VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html)
    that selects between using state specified by that structure and
    state specified by
    [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html).

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2020-09-02 (pbrown)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_fragment_shading_rate_enums"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
