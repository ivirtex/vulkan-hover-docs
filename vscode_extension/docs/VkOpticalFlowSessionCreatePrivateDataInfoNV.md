# VkOpticalFlowSessionCreatePrivateDataInfoNV(3) Manual Page

## Name

VkOpticalFlowSessionCreatePrivateDataInfoNV - Structure for NV internal use only



## [](#_c_specification)C Specification

The [VkOpticalFlowSessionCreatePrivateDataInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionCreatePrivateDataInfoNV.html) structure is for NV internal use only and is defined as:

```c++
// Provided by VK_NV_optical_flow
typedef struct VkOpticalFlowSessionCreatePrivateDataInfoNV {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           id;
    uint32_t           size;
    const void*        pPrivateData;
} VkOpticalFlowSessionCreatePrivateDataInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `id` is an identifier for data which is passed at a memory location specified in `VkOpticalFlowSessionCreatePrivateDataInfoNV`::`pPrivateData`.
- `size` is the size of data in bytes which is passed at a memory location specified in `VkOpticalFlowSessionCreatePrivateDataInfoNV`::`pPrivateData`.
- `pPrivateData` is a pointer to NV internal data.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkOpticalFlowSessionCreatePrivateDataInfoNV-sType-sType)VUID-VkOpticalFlowSessionCreatePrivateDataInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_OPTICAL_FLOW_SESSION_CREATE_PRIVATE_DATA_INFO_NV`
- [](#VUID-VkOpticalFlowSessionCreatePrivateDataInfoNV-pPrivateData-parameter)VUID-VkOpticalFlowSessionCreatePrivateDataInfoNV-pPrivateData-parameter  
  `pPrivateData` **must** be a pointer value

## [](#_see_also)See Also

[VK\_NV\_optical\_flow](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_optical_flow.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkOpticalFlowSessionCreatePrivateDataInfoNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0