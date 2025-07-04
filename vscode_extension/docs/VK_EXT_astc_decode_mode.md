# VK\_EXT\_astc\_decode\_mode(3) Manual Page

## Name

VK\_EXT\_astc\_decode\_mode - device extension



## [](#_registered_extension_number)Registered Extension Number

68

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Jan-Harald Fredriksen [\[GitHub\]janharaldfredriksen-arm](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_astc_decode_mode%5D%20%40janharaldfredriksen-arm%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_astc_decode_mode%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-08-07

**Contributors**

- Jan-Harald Fredriksen, Arm

## [](#_description)Description

The existing specification requires that low dynamic range (LDR) ASTC textures are decompressed to FP16 values per component. In many cases, decompressing LDR textures to a lower precision intermediate result gives acceptable image quality. Source material for LDR textures is typically authored as 8-bit UNORM values, so decoding to FP16 values adds little value. On the other hand, reducing precision of the decoded result reduces the size of the decompressed data, potentially improving texture cache performance and saving power.

The goal of this extension is to enable this efficiency gain on existing ASTC texture data. This is achieved by giving the application the ability to select the intermediate decoding precision.

Three decoding options are provided:

- Decode to `VK_FORMAT_R16G16B16A16_SFLOAT` precision: This is the default, and matches the required behavior in the core API.
- Decode to `VK_FORMAT_R8G8B8A8_UNORM` precision: This is provided as an option in LDR mode.
- Decode to `VK_FORMAT_E5B9G9R9_UFLOAT_PACK32` precision: This is provided as an option in both LDR and HDR mode. In this mode, negative values cannot be represented and are clamped to zero. The alpha component is ignored, and the results are as if alpha was 1.0. This decode mode is optional and support can be queried via the physical device properties.

## [](#_new_structures)New Structures

- Extending [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html):
  
  - [VkImageViewASTCDecodeModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewASTCDecodeModeEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceASTCDecodeFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceASTCDecodeFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_ASTC_DECODE_MODE_EXTENSION_NAME`
- `VK_EXT_ASTC_DECODE_MODE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_IMAGE_VIEW_ASTC_DECODE_MODE_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ASTC_DECODE_FEATURES_EXT`

## [](#_issues)Issues

1\) Are implementations allowed to decode at a higher precision than what is requested?

```
RESOLUTION: No.
If we allow this, then this extension could be exposed on all
implementations that support ASTC.
But developers would have no way of knowing what precision was actually
used, and thus whether the image quality is sufficient at reduced
precision.
```

2\) Should the decode mode be image view state and/or sampler state?

```
RESOLUTION: Image view state only.
Some implementations treat the different decode modes as different
texture formats.
```

## [](#_example)Example

Create an image view that decodes to `VK_FORMAT_R8G8B8A8_UNORM` precision:

```c++
    VkImageViewASTCDecodeModeEXT decodeMode =
    {
        .sType = VK_STRUCTURE_TYPE_IMAGE_VIEW_ASTC_DECODE_MODE_EXT,
        .pNext = NULL,
        .decodeMode = VK_FORMAT_R8G8B8A8_UNORM
    };

    VkImageViewCreateInfo createInfo =
    {
        .sType = VK_STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO,
        .pNext = &decodeMode,
        // flags, image, viewType set to application-desired values
        .format = VK_FORMAT_ASTC_8x8_UNORM_BLOCK,
        // components, subresourceRange set to application-desired values
    };

    VkImageView imageView;
    VkResult result = vkCreateImageView(
        device,
        &createInfo,
        NULL,
        &imageView);
```

## [](#_version_history)Version History

- Revision 1, 2018-08-07 (Jan-Harald Fredriksen)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_astc_decode_mode)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0