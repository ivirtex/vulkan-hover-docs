# VkAndroidHardwareBufferFormatResolvePropertiesANDROID(3) Manual Page

## Name

VkAndroidHardwareBufferFormatResolvePropertiesANDROID - Structure
defining properties of resolves using an external format



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkAndroidHardwareBufferFormatResolvePropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferFormatResolvePropertiesANDROID.html)
structure is defined as:

``` c
// Provided by VK_ANDROID_external_format_resolve
typedef struct VkAndroidHardwareBufferFormatResolvePropertiesANDROID {
    VkStructureType    sType;
    void*              pNext;
    VkFormat           colorAttachmentFormat;
} VkAndroidHardwareBufferFormatResolvePropertiesANDROID;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `colorAttachmentFormat` is a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) specifying the
  format of color attachment images that **must** be used for color
  attachments when resolving to the specified external format. If the
  implementation supports external format resolves for the specified
  external format, this value will be set to a color format supporting
  the `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT` in
  [VkFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties.html)::`optimalTilingFeatures`
  as returned by
  [vkGetPhysicalDeviceFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties.html)
  with `format` equal to `colorAttachmentFormat` If external format
  resolves are not supported, this value will be set to
  `VK_FORMAT_UNDEFINED`.

## <a href="#_description" class="anchor"></a>Description

Any Android hardware buffer created with the `GRALLOC_USAGE_HW_RENDER`
flag **must** be renderable in some way in Vulkan, either:

- [VkAndroidHardwareBufferFormatPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferFormatPropertiesANDROID.html)::`format`
  **must** be a format that supports
  `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT` or
  `VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT` in
  [VkFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties.html)::`optimalTilingFeatures`;
  or

- `colorAttachmentFormat` **must** be a format that supports
  `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT` in
  [VkFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties.html)::`optimalTilingFeatures`.

Valid Usage (Implicit)

- <a
  href="#VUID-VkAndroidHardwareBufferFormatResolvePropertiesANDROID-sType-sType"
  id="VUID-VkAndroidHardwareBufferFormatResolvePropertiesANDROID-sType-sType"></a>
  VUID-VkAndroidHardwareBufferFormatResolvePropertiesANDROID-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_FORMAT_RESOLVE_PROPERTIES_ANDROID`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_ANDROID_external_format_resolve](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_ANDROID_external_format_resolve.html),
[VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAndroidHardwareBufferFormatResolvePropertiesANDROID"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
