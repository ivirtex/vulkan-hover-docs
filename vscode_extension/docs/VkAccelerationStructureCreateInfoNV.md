# VkAccelerationStructureCreateInfoNV(3) Manual Page

## Name

VkAccelerationStructureCreateInfoNV - Structure specifying the parameters of a newly created acceleration structure object



## [](#_c_specification)C Specification

The `VkAccelerationStructureCreateInfoNV` structure is defined as:

```c++
// Provided by VK_NV_ray_tracing
typedef struct VkAccelerationStructureCreateInfoNV {
    VkStructureType                  sType;
    const void*                      pNext;
    VkDeviceSize                     compactedSize;
    VkAccelerationStructureInfoNV    info;
} VkAccelerationStructureCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `compactedSize` is the size from the result of [vkCmdWriteAccelerationStructuresPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteAccelerationStructuresPropertiesNV.html) if this acceleration structure is going to be the target of a compacting copy.
- `info` is the [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInfoNV.html) structure specifying further parameters of the created acceleration structure.

## [](#_description)Description

Valid Usage

- [](#VUID-VkAccelerationStructureCreateInfoNV-compactedSize-02421)VUID-VkAccelerationStructureCreateInfoNV-compactedSize-02421  
  If `compactedSize` is not `0` then both `info.geometryCount` and `info.instanceCount` **must** be `0`

Valid Usage (Implicit)

- [](#VUID-VkAccelerationStructureCreateInfoNV-sType-sType)VUID-VkAccelerationStructureCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_CREATE_INFO_NV`
- [](#VUID-VkAccelerationStructureCreateInfoNV-pNext-pNext)VUID-VkAccelerationStructureCreateInfoNV-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html)
- [](#VUID-VkAccelerationStructureCreateInfoNV-sType-unique)VUID-VkAccelerationStructureCreateInfoNV-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkAccelerationStructureCreateInfoNV-info-parameter)VUID-VkAccelerationStructureCreateInfoNV-info-parameter  
  `info` **must** be a valid [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInfoNV.html) structure

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInfoNV.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateAccelerationStructureNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureCreateInfoNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0