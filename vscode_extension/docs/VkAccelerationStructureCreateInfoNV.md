# VkAccelerationStructureCreateInfoNV(3) Manual Page

## Name

VkAccelerationStructureCreateInfoNV - Structure specifying the
parameters of a newly created acceleration structure object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAccelerationStructureCreateInfoNV` structure is defined as:

``` c
// Provided by VK_NV_ray_tracing
typedef struct VkAccelerationStructureCreateInfoNV {
    VkStructureType                  sType;
    const void*                      pNext;
    VkDeviceSize                     compactedSize;
    VkAccelerationStructureInfoNV    info;
} VkAccelerationStructureCreateInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `compactedSize` is the size from the result of
  [vkCmdWriteAccelerationStructuresPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteAccelerationStructuresPropertiesNV.html)
  if this acceleration structure is going to be the target of a
  compacting copy.

- `info` is the
  [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInfoNV.html)
  structure specifying further parameters of the created acceleration
  structure.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkAccelerationStructureCreateInfoNV-compactedSize-02421"
  id="VUID-VkAccelerationStructureCreateInfoNV-compactedSize-02421"></a>
  VUID-VkAccelerationStructureCreateInfoNV-compactedSize-02421  
  If `compactedSize` is not `0` then both `info.geometryCount` and
  `info.instanceCount` **must** be `0`

Valid Usage (Implicit)

- <a href="#VUID-VkAccelerationStructureCreateInfoNV-sType-sType"
  id="VUID-VkAccelerationStructureCreateInfoNV-sType-sType"></a>
  VUID-VkAccelerationStructureCreateInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_CREATE_INFO_NV`

- <a href="#VUID-VkAccelerationStructureCreateInfoNV-pNext-pNext"
  id="VUID-VkAccelerationStructureCreateInfoNV-pNext-pNext"></a>
  VUID-VkAccelerationStructureCreateInfoNV-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html)

- <a href="#VUID-VkAccelerationStructureCreateInfoNV-sType-unique"
  id="VUID-VkAccelerationStructureCreateInfoNV-sType-unique"></a>
  VUID-VkAccelerationStructureCreateInfoNV-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkAccelerationStructureCreateInfoNV-info-parameter"
  id="VUID-VkAccelerationStructureCreateInfoNV-info-parameter"></a>
  VUID-VkAccelerationStructureCreateInfoNV-info-parameter  
  `info` **must** be a valid
  [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInfoNV.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInfoNV.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateAccelerationStructureNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAccelerationStructureCreateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
