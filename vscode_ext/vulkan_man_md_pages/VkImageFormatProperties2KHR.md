# VkImageFormatProperties2(3) Manual Page

## Name

VkImageFormatProperties2 - Structure specifying an image format
properties



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImageFormatProperties2` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkImageFormatProperties2 {
    VkStructureType            sType;
    void*                      pNext;
    VkImageFormatProperties    imageFormatProperties;
} VkImageFormatProperties2;
```

or the equivalent

``` c
// Provided by VK_KHR_get_physical_device_properties2
typedef VkImageFormatProperties2 VkImageFormatProperties2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure. The `pNext` chain of `VkImageFormatProperties2` is used to
  allow the specification of additional capabilities to be returned from
  `vkGetPhysicalDeviceImageFormatProperties2`.

- `imageFormatProperties` is a
  [VkImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties.html) structure in
  which capabilities are returned.

## <a href="#_description" class="anchor"></a>Description

If the combination of parameters to
`vkGetPhysicalDeviceImageFormatProperties2` is not supported by the
implementation for use in [vkCreateImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImage.html), then all
members of `imageFormatProperties` will be filled with zero.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Filling <code>imageFormatProperties</code> with zero for unsupported
formats is an exception to the usual rule that output structures have
undefined contents on error. This exception was unintentional, but is
preserved for backwards compatibility. This exception only applies to
<code>imageFormatProperties</code>, not <code>sType</code>,
<code>pNext</code>, or any structures chained from
<code>pNext</code>.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-VkImageFormatProperties2-sType-sType"
  id="VUID-VkImageFormatProperties2-sType-sType"></a>
  VUID-VkImageFormatProperties2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2`

- <a href="#VUID-VkImageFormatProperties2-pNext-pNext"
  id="VUID-VkImageFormatProperties2-pNext-pNext"></a>
  VUID-VkImageFormatProperties2-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkAndroidHardwareBufferUsageANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferUsageANDROID.html),
  [VkExternalImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalImageFormatProperties.html),
  [VkFilterCubicImageViewImageFormatPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFilterCubicImageViewImageFormatPropertiesEXT.html),
  [VkHostImageCopyDevicePerformanceQueryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHostImageCopyDevicePerformanceQueryEXT.html),
  [VkImageCompressionPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionPropertiesEXT.html),
  [VkSamplerYcbcrConversionImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionImageFormatProperties.html),
  or
  [VkTextureLODGatherFormatPropertiesAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTextureLODGatherFormatPropertiesAMD.html)

- <a href="#VUID-VkImageFormatProperties2-sType-unique"
  id="VUID-VkImageFormatProperties2-sType-unique"></a>
  VUID-VkImageFormatProperties2-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html),
[vkGetPhysicalDeviceImageFormatProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageFormatProperties2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
