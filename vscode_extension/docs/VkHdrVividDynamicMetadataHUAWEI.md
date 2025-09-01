# VkHdrVividDynamicMetadataHUAWEI(3) Manual Page

## Name

VkHdrVividDynamicMetadataHUAWEI - specify HDR Vivid dynamic metadata



## [](#_c_specification)C Specification

When [`hdrVivid`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-hdrVivid) feature is enabled, HDR Vivid dynamic metadata **can** be set to control the reproduction of content by including the `VkHdrVividDynamicMetadataHUAWEI` in the `pNext` chain of [VkHdrMetadataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHdrMetadataEXT.html).

The `VkHdrVividDynamicMetadataHUAWEI` structure is defined as:

```c++
// Provided by VK_HUAWEI_hdr_vivid
typedef struct VkHdrVividDynamicMetadataHUAWEI {
    VkStructureType    sType;
    const void*        pNext;
    size_t             dynamicMetadataSize;
    const void*        pDynamicMetadata;
} VkHdrVividDynamicMetadataHUAWEI;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `dynamicMetadataSize` is the size in bytes of the dynamic metadata.
- `pDynamicMetadata` is a pointer to the dynamic metadata.

## [](#_description)Description

Note

The HDR Vivid metadata is intended to be used as defined in the T/UWA 005.1-2022 specification. The validity and use of this data is outside the scope of Vulkan.

Valid Usage (Implicit)

- [](#VUID-VkHdrVividDynamicMetadataHUAWEI-sType-sType)VUID-VkHdrVividDynamicMetadataHUAWEI-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_HDR_VIVID_DYNAMIC_METADATA_HUAWEI`
- [](#VUID-VkHdrVividDynamicMetadataHUAWEI-pDynamicMetadata-parameter)VUID-VkHdrVividDynamicMetadataHUAWEI-pDynamicMetadata-parameter  
  `pDynamicMetadata` **must** be a valid pointer to an array of `dynamicMetadataSize` bytes
- [](#VUID-VkHdrVividDynamicMetadataHUAWEI-dynamicMetadataSize-arraylength)VUID-VkHdrVividDynamicMetadataHUAWEI-dynamicMetadataSize-arraylength  
  `dynamicMetadataSize` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_HUAWEI\_hdr\_vivid](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_HUAWEI_hdr_vivid.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkHdrVividDynamicMetadataHUAWEI).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0