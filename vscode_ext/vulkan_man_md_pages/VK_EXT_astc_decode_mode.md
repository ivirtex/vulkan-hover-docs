# VK_EXT_astc_decode_mode(3) Manual Page

## Name

VK_EXT_astc_decode_mode - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

68

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Jan-Harald Fredriksen <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_astc_decode_mode%5D%20@janharaldfredriksen-arm%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_astc_decode_mode%20extension*"
  target="_blank"
  rel="nofollow noopener"><em></em>janharaldfredriksen-arm</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-08-07

**Contributors**  
- Jan-Harald Fredriksen, Arm

## <a href="#_description" class="anchor"></a>Description

The existing specification requires that low dynamic range (LDR) ASTC
textures are decompressed to FP16 values per component. In many cases,
decompressing LDR textures to a lower precision intermediate result
gives acceptable image quality. Source material for LDR textures is
typically authored as 8-bit UNORM values, so decoding to FP16 values
adds little value. On the other hand, reducing precision of the decoded
result reduces the size of the decompressed data, potentially improving
texture cache performance and saving power.

The goal of this extension is to enable this efficiency gain on existing
ASTC texture data. This is achieved by giving the application the
ability to select the intermediate decoding precision.

Three decoding options are provided:

- Decode to `VK_FORMAT_R16G16B16A16_SFLOAT` precision: This is the
  default, and matches the required behavior in the core API.

- Decode to `VK_FORMAT_R8G8B8A8_UNORM` precision: This is provided as an
  option in LDR mode.

- Decode to `VK_FORMAT_E5B9G9R9_UFLOAT_PACK32` precision: This is
  provided as an option in both LDR and HDR mode. In this mode, negative
  values cannot be represented and are clamped to zero. The alpha
  component is ignored, and the results are as if alpha was 1.0. This
  decode mode is optional and support can be queried via the physical
  device properties.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html):

  - [VkImageViewASTCDecodeModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewASTCDecodeModeEXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceASTCDecodeFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceASTCDecodeFeaturesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_ASTC_DECODE_MODE_EXTENSION_NAME`

- `VK_EXT_ASTC_DECODE_MODE_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_IMAGE_VIEW_ASTC_DECODE_MODE_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ASTC_DECODE_FEATURES_EXT`

## <a href="#_issues" class="anchor"></a>Issues

1\) Are implementations allowed to decode at a higher precision than
what is requested?

    RESOLUTION: No.
    If we allow this, then this extension could be exposed on all
    implementations that support ASTC.
    But developers would have no way of knowing what precision was actually
    used, and thus whether the image quality is sufficient at reduced
    precision.

2\) Should the decode mode be image view state and/or sampler state?

    RESOLUTION: Image view state only.
    Some implementations treat the different decode modes as different
    texture formats.

## <a href="#_example" class="anchor"></a>Example

Create an image view that decodes to `VK_FORMAT_R8G8B8A8_UNORM`
precision:

``` c
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

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-08-07 (Jan-Harald Fredriksen)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_astc_decode_mode"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
