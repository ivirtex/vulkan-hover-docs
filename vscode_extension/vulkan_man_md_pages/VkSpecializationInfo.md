# VkSpecializationInfo(3) Manual Page

## Name

VkSpecializationInfo - Structure specifying specialization information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSpecializationInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkSpecializationInfo {
    uint32_t                           mapEntryCount;
    const VkSpecializationMapEntry*    pMapEntries;
    size_t                             dataSize;
    const void*                        pData;
} VkSpecializationInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `mapEntryCount` is the number of entries in the `pMapEntries` array.

- `pMapEntries` is a pointer to an array of `VkSpecializationMapEntry`
  structures, which map constant IDs to offsets in `pData`.

- `dataSize` is the byte size of the `pData` buffer.

- `pData` contains the actual constant values to specialize with.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkSpecializationInfo-offset-00773"
  id="VUID-VkSpecializationInfo-offset-00773"></a>
  VUID-VkSpecializationInfo-offset-00773  
  The `offset` member of each element of `pMapEntries` **must** be less
  than `dataSize`

- <a href="#VUID-VkSpecializationInfo-pMapEntries-00774"
  id="VUID-VkSpecializationInfo-pMapEntries-00774"></a>
  VUID-VkSpecializationInfo-pMapEntries-00774  
  The `size` member of each element of `pMapEntries` **must** be less
  than or equal to `dataSize` minus `offset`

- <a href="#VUID-VkSpecializationInfo-constantID-04911"
  id="VUID-VkSpecializationInfo-constantID-04911"></a>
  VUID-VkSpecializationInfo-constantID-04911  
  The `constantID` value of each element of `pMapEntries` **must** be
  unique within `pMapEntries`

Valid Usage (Implicit)

- <a href="#VUID-VkSpecializationInfo-pMapEntries-parameter"
  id="VUID-VkSpecializationInfo-pMapEntries-parameter"></a>
  VUID-VkSpecializationInfo-pMapEntries-parameter  
  If `mapEntryCount` is not `0`, `pMapEntries` **must** be a valid
  pointer to an array of `mapEntryCount` valid
  [VkSpecializationMapEntry](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSpecializationMapEntry.html) structures

- <a href="#VUID-VkSpecializationInfo-pData-parameter"
  id="VUID-VkSpecializationInfo-pData-parameter"></a>
  VUID-VkSpecializationInfo-pData-parameter  
  If `dataSize` is not `0`, `pData` **must** be a valid pointer to an
  array of `dataSize` bytes

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html),
[VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateInfoEXT.html),
[VkSpecializationMapEntry](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSpecializationMapEntry.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSpecializationInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
