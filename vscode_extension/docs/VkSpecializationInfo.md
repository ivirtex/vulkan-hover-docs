# VkSpecializationInfo(3) Manual Page

## Name

VkSpecializationInfo - Structure specifying specialization information



## [](#_c_specification)C Specification

The `VkSpecializationInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkSpecializationInfo {
    uint32_t                           mapEntryCount;
    const VkSpecializationMapEntry*    pMapEntries;
    size_t                             dataSize;
    const void*                        pData;
} VkSpecializationInfo;
```

## [](#_members)Members

- `mapEntryCount` is the number of entries in the `pMapEntries` array.
- `pMapEntries` is a pointer to an array of `VkSpecializationMapEntry` structures, which map constant IDs to offsets in `pData`.
- `dataSize` is the byte size of the `pData` buffer.
- `pData` contains the actual constant values to specialize with.

## [](#_description)Description

Valid Usage

- [](#VUID-VkSpecializationInfo-offset-00773)VUID-VkSpecializationInfo-offset-00773  
  The `offset` member of each element of `pMapEntries` **must** be less than `dataSize`
- [](#VUID-VkSpecializationInfo-pMapEntries-00774)VUID-VkSpecializationInfo-pMapEntries-00774  
  The `size` member of each element of `pMapEntries` **must** be less than or equal to `dataSize` minus `offset`
- [](#VUID-VkSpecializationInfo-constantID-04911)VUID-VkSpecializationInfo-constantID-04911  
  The `constantID` value of each element of `pMapEntries` **must** be unique within `pMapEntries`

Valid Usage (Implicit)

- [](#VUID-VkSpecializationInfo-pMapEntries-parameter)VUID-VkSpecializationInfo-pMapEntries-parameter  
  If `mapEntryCount` is not `0`, `pMapEntries` **must** be a valid pointer to an array of `mapEntryCount` valid [VkSpecializationMapEntry](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSpecializationMapEntry.html) structures
- [](#VUID-VkSpecializationInfo-pData-parameter)VUID-VkSpecializationInfo-pData-parameter  
  If `dataSize` is not `0`, `pData` **must** be a valid pointer to an array of `dataSize` bytes

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDataGraphPipelineShaderModuleCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineShaderModuleCreateInfoARM.html), [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html), [VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCreateInfoEXT.html), [VkSpecializationMapEntry](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSpecializationMapEntry.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSpecializationInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0