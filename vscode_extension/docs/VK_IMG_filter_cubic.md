# VK\_IMG\_filter\_cubic(3) Manual Page

## Name

VK\_IMG\_filter\_cubic - device extension



## [](#_registered_extension_number)Registered Extension Number

16

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- Tobias Hector [\[GitHub\]tobski](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_IMG_filter_cubic%5D%20%40tobski%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_IMG_filter_cubic%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2016-02-23

**Contributors**

- Tobias Hector, Imagination Technologies

## [](#_description)Description

`VK_IMG_filter_cubic` adds an additional, high quality cubic filtering mode to Vulkan, using a Catmull-Rom bicubic filter. Performing this kind of filtering can be done in a shader by using 16 samples and a number of instructions, but this can be inefficient. The cubic filter mode exposes an optimized high quality texture sampling using fixed texture sampling functionality.

## [](#_new_enum_constants)New Enum Constants

- `VK_IMG_FILTER_CUBIC_EXTENSION_NAME`
- `VK_IMG_FILTER_CUBIC_SPEC_VERSION`
- Extending [VkFilter](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFilter.html):
  
  - `VK_FILTER_CUBIC_IMG`
- Extending [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits.html):
  
  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_CUBIC_BIT_IMG`

## [](#_example)Example

Creating a sampler with the new filter for both magnification and minification

```c++
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

## [](#_version_history)Version History

- Revision 1, 2016-02-23 (Tobias Hector)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_IMG_filter_cubic).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0