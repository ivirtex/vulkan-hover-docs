# VK_IMG_filter_cubic(3) Manual Page

## Name

VK_IMG_filter_cubic - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

16

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- Tobias Hector <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_IMG_filter_cubic%5D%20@tobski%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_IMG_filter_cubic%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>tobski</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2016-02-23

**Contributors**  
- Tobias Hector, Imagination Technologies

## <a href="#_description" class="anchor"></a>Description

`VK_IMG_filter_cubic` adds an additional, high quality cubic filtering
mode to Vulkan, using a Catmull-Rom bicubic filter. Performing this kind
of filtering can be done in a shader by using 16 samples and a number of
instructions, but this can be inefficient. The cubic filter mode exposes
an optimized high quality texture sampling using fixed texture sampling
functionality.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_IMG_FILTER_CUBIC_EXTENSION_NAME`

- `VK_IMG_FILTER_CUBIC_SPEC_VERSION`

- Extending [VkFilter](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFilter.html):

  - `VK_FILTER_CUBIC_IMG`

- Extending [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html):

  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_IMG`

## <a href="#_example" class="anchor"></a>Example

Creating a sampler with the new filter for both magnification and
minification

``` c
    VkSamplerCreateInfo createInfo =
    {
        .sType = VK_STRUCTURE_TYPE_SAMPLER_CREATE_INFO,
        // Other members set to application-desired values
    };

    createInfo.magFilter = VK_FILTER_CUBIC_IMG;
    createInfo.minFilter = VK_FILTER_CUBIC_IMG;

    VkSampler sampler;
    VkResult result = vkCreateSampler(
        device,
        &createInfo,
        &sampler);
```

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-02-23 (Tobias Hector)

  - Initial version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_IMG_filter_cubic"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
