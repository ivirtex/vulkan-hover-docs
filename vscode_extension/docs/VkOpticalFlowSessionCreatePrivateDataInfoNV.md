# VkOpticalFlowSessionCreatePrivateDataInfoNV(3) Manual Page

## Name

VkOpticalFlowSessionCreatePrivateDataInfoNV - Structure for NV internal
use only



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkOpticalFlowSessionCreatePrivateDataInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionCreatePrivateDataInfoNV.html)
structure is for NV internal use only and is defined as:

``` c
// Provided by VK_NV_optical_flow
typedef struct VkOpticalFlowSessionCreatePrivateDataInfoNV {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           id;
    uint32_t           size;
    const void*        pPrivateData;
} VkOpticalFlowSessionCreatePrivateDataInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `id` is an identifier for data which is passed at a memory location
  specified in
  `VkOpticalFlowSessionCreatePrivateDataInfoNV`::`pPrivateData`.

- `size` is the size of data in bytes which is passed at a memory
  location specified in
  `VkOpticalFlowSessionCreatePrivateDataInfoNV`::`pPrivateData`.

- `pPrivateData` is a pointer to NV internal data.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkOpticalFlowSessionCreatePrivateDataInfoNV-sType-sType"
  id="VUID-VkOpticalFlowSessionCreatePrivateDataInfoNV-sType-sType"></a>
  VUID-VkOpticalFlowSessionCreatePrivateDataInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_OPTICAL_FLOW_SESSION_CREATE_PRIVATE_DATA_INFO_NV`

- <a
  href="#VUID-VkOpticalFlowSessionCreatePrivateDataInfoNV-pPrivateData-parameter"
  id="VUID-VkOpticalFlowSessionCreatePrivateDataInfoNV-pPrivateData-parameter"></a>
  VUID-VkOpticalFlowSessionCreatePrivateDataInfoNV-pPrivateData-parameter  
  `pPrivateData` **must** be a pointer value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_optical_flow](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_optical_flow.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkOpticalFlowSessionCreatePrivateDataInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
