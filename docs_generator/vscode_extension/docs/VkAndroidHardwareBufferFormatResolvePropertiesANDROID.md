# VkAndroidHardwareBufferFormatResolvePropertiesANDROID(3) Manual Page

## Name

VkAndroidHardwareBufferFormatResolvePropertiesANDROID - Structure defining properties of resolves using an external format



## [](#_c_specification)C Specification

The [VkAndroidHardwareBufferFormatResolvePropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferFormatResolvePropertiesANDROID.html) structure is defined as:

```c++
// Provided by VK_ANDROID_external_format_resolve
typedef struct VkAndroidHardwareBufferFormatResolvePropertiesANDROID {
    VkStructureType    sType;
    void*              pNext;
    VkFormat           colorAttachmentFormat;
} VkAndroidHardwareBufferFormatResolvePropertiesANDROID;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `colorAttachmentFormat` is a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) specifying the format of color attachment images that **must** be used for color attachments when resolving to the specified external format. If the implementation supports external format resolves for the specified external format, this value will be a color format supporting the `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT` in [VkFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties.html)::`optimalTilingFeatures` as returned by [vkGetPhysicalDeviceFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFormatProperties.html) with `format` equal to `colorAttachmentFormat` If external format resolves are not supported, this value will be `VK_FORMAT_UNDEFINED`.

## [](#_description)Description

Any Android hardware buffer created with the `GRALLOC_USAGE_HW_RENDER` flag **must** be renderable in some way in Vulkan, either:

- [VkAndroidHardwareBufferFormatPropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferFormatPropertiesANDROID.html)::`format` **must** be a format that supports `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT` or `VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT` in [VkFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties.html)::`optimalTilingFeatures`; or
- `colorAttachmentFormat` **must** be a format that supports `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT` in [VkFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties.html)::`optimalTilingFeatures`.

Valid Usage (Implicit)

- [](#VUID-VkAndroidHardwareBufferFormatResolvePropertiesANDROID-sType-sType)VUID-VkAndroidHardwareBufferFormatResolvePropertiesANDROID-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_FORMAT_RESOLVE_PROPERTIES_ANDROID`

## [](#_see_also)See Also

[VK\_ANDROID\_external\_format\_resolve](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ANDROID_external_format_resolve.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAndroidHardwareBufferFormatResolvePropertiesANDROID)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0