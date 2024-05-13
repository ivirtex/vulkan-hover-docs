# VkOpaqueCaptureDescriptorDataCreateInfoEXT(3) Manual Page

## Name

VkOpaqueCaptureDescriptorDataCreateInfoEXT - Structure specifying opaque
capture descriptor data



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkOpaqueCaptureDescriptorDataCreateInfoEXT` structure is defined
as:

``` c
// Provided by VK_EXT_descriptor_buffer
typedef struct VkOpaqueCaptureDescriptorDataCreateInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    const void*        opaqueCaptureDescriptorData;
} VkOpaqueCaptureDescriptorDataCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `opaqueCaptureDescriptorData` is a pointer to a user-allocated buffer
  containing opaque capture data retrieved using
  [vkGetBufferOpaqueCaptureDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferOpaqueCaptureDescriptorDataEXT.html),
  [vkGetImageOpaqueCaptureDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageOpaqueCaptureDescriptorDataEXT.html),
  [vkGetImageViewOpaqueCaptureDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageViewOpaqueCaptureDescriptorDataEXT.html),
  [vkGetSamplerOpaqueCaptureDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetSamplerOpaqueCaptureDescriptorDataEXT.html),
  or
  [vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT.html).

## <a href="#_description" class="anchor"></a>Description

During replay, opaque descriptor capture data **can** be specified by
adding a `VkOpaqueCaptureDescriptorDataCreateInfoEXT` structure to the
relevant `pNext` chain of a
[VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html),
[VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html),
[VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html),
[VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html),
[VkAccelerationStructureCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoNV.html)
or
[VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoKHR.html)
structure.

Valid Usage (Implicit)

- <a href="#VUID-VkOpaqueCaptureDescriptorDataCreateInfoEXT-sType-sType"
  id="VUID-VkOpaqueCaptureDescriptorDataCreateInfoEXT-sType-sType"></a>
  VUID-VkOpaqueCaptureDescriptorDataCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_OPAQUE_CAPTURE_DESCRIPTOR_DATA_CREATE_INFO_EXT`

- <a
  href="#VUID-VkOpaqueCaptureDescriptorDataCreateInfoEXT-opaqueCaptureDescriptorData-parameter"
  id="VUID-VkOpaqueCaptureDescriptorDataCreateInfoEXT-opaqueCaptureDescriptorData-parameter"></a>
  VUID-VkOpaqueCaptureDescriptorDataCreateInfoEXT-opaqueCaptureDescriptorData-parameter  
  `opaqueCaptureDescriptorData` **must** be a pointer value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkOpaqueCaptureDescriptorDataCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
