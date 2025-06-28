# vkWriteMicromapsPropertiesEXT(3) Manual Page

## Name

vkWriteMicromapsPropertiesEXT - Query micromap meta-data on the host



## [](#_c_specification)C Specification

To query micromap size parameters on the host, call:

```c++
// Provided by VK_EXT_opacity_micromap
VkResult vkWriteMicromapsPropertiesEXT(
    VkDevice                                    device,
    uint32_t                                    micromapCount,
    const VkMicromapEXT*                        pMicromaps,
    VkQueryType                                 queryType,
    size_t                                      dataSize,
    void*                                       pData,
    size_t                                      stride);
```

## [](#_parameters)Parameters

- `device` is the device which owns the micromaps in `pMicromaps`.
- `micromapCount` is the count of micromaps for which to query the property.
- `pMicromaps` is a pointer to an array of existing previously built micromaps.
- `queryType` is a [VkQueryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryType.html) value specifying the property to be queried.
- `dataSize` is the size in bytes of the buffer pointed to by `pData`.
- `pData` is a pointer to an application-allocated buffer where the results will be written.
- `stride` is the stride in bytes between results for individual queries within `pData`.

## [](#_description)Description

This command fulfills the same task as [vkCmdWriteMicromapsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteMicromapsPropertiesEXT.html) but is executed by the host.

Valid Usage

- [](#VUID-vkWriteMicromapsPropertiesEXT-pMicromaps-07501)VUID-vkWriteMicromapsPropertiesEXT-pMicromaps-07501  
  All micromaps in `pMicromaps` **must** have been constructed prior to the execution of this command
- [](#VUID-vkWriteMicromapsPropertiesEXT-pMicromaps-07502)VUID-vkWriteMicromapsPropertiesEXT-pMicromaps-07502  
  All micromaps in `pMicromaps` **must** have been constructed with `VK_BUILD_MICROMAP_ALLOW_COMPACTION_BIT_EXT` if `queryType` is `VK_QUERY_TYPE_MICROMAP_COMPACTED_SIZE_EXT`
- [](#VUID-vkWriteMicromapsPropertiesEXT-queryType-07503)VUID-vkWriteMicromapsPropertiesEXT-queryType-07503  
  `queryType` **must** be `VK_QUERY_TYPE_MICROMAP_COMPACTED_SIZE_EXT` or `VK_QUERY_TYPE_MICROMAP_SERIALIZATION_SIZE_EXT`
- [](#VUID-vkWriteMicromapsPropertiesEXT-queryType-10071)VUID-vkWriteMicromapsPropertiesEXT-queryType-10071  
  If `queryType` is `VK_QUERY_TYPE_MICROMAP_SERIALIZATION_SIZE_EXT` or `VK_QUERY_TYPE_MICROMAP_COMPACTED_SIZE_EXT` then `stride` **must** be a multiple of the size of [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)
- [](#VUID-vkWriteMicromapsPropertiesEXT-queryType-10072)VUID-vkWriteMicromapsPropertiesEXT-queryType-10072  
  If `queryType` is `VK_QUERY_TYPE_MICROMAP_SERIALIZATION_SIZE_EXT` or `VK_QUERY_TYPE_MICROMAP_COMPACTED_SIZE_EXT` then `pData` **must** point to a [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)
- [](#VUID-vkWriteMicromapsPropertiesEXT-dataSize-07576)VUID-vkWriteMicromapsPropertiesEXT-dataSize-07576  
  `dataSize` **must** be greater than or equal to `micromapCount`\*`stride`
- [](#VUID-vkWriteMicromapsPropertiesEXT-buffer-07577)VUID-vkWriteMicromapsPropertiesEXT-buffer-07577  
  The `buffer` used to create each micromap in `pMicromaps` **must** be bound to host-visible device memory
- [](#VUID-vkWriteMicromapsPropertiesEXT-micromapHostCommands-07578)VUID-vkWriteMicromapsPropertiesEXT-micromapHostCommands-07578  
  The [`VkPhysicalDeviceOpacityMicromapFeaturesEXT`::`micromapHostCommands`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-micromapHostCommands) feature **must** be enabled
- [](#VUID-vkWriteMicromapsPropertiesEXT-buffer-07579)VUID-vkWriteMicromapsPropertiesEXT-buffer-07579  
  The `buffer` used to create each micromap in `pMicromaps` **must** be bound to memory that was not allocated with multiple instances

Valid Usage (Implicit)

- [](#VUID-vkWriteMicromapsPropertiesEXT-device-parameter)VUID-vkWriteMicromapsPropertiesEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkWriteMicromapsPropertiesEXT-pMicromaps-parameter)VUID-vkWriteMicromapsPropertiesEXT-pMicromaps-parameter  
  `pMicromaps` **must** be a valid pointer to an array of `micromapCount` valid [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapEXT.html) handles
- [](#VUID-vkWriteMicromapsPropertiesEXT-queryType-parameter)VUID-vkWriteMicromapsPropertiesEXT-queryType-parameter  
  `queryType` **must** be a valid [VkQueryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryType.html) value
- [](#VUID-vkWriteMicromapsPropertiesEXT-pData-parameter)VUID-vkWriteMicromapsPropertiesEXT-pData-parameter  
  `pData` **must** be a valid pointer to an array of `dataSize` bytes
- [](#VUID-vkWriteMicromapsPropertiesEXT-micromapCount-arraylength)VUID-vkWriteMicromapsPropertiesEXT-micromapCount-arraylength  
  `micromapCount` **must** be greater than `0`
- [](#VUID-vkWriteMicromapsPropertiesEXT-dataSize-arraylength)VUID-vkWriteMicromapsPropertiesEXT-dataSize-arraylength  
  `dataSize` **must** be greater than `0`
- [](#VUID-vkWriteMicromapsPropertiesEXT-pMicromaps-parent)VUID-vkWriteMicromapsPropertiesEXT-pMicromaps-parent  
  Each element of `pMicromaps` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapEXT.html), [VkQueryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkWriteMicromapsPropertiesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0