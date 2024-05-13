# vkWriteMicromapsPropertiesEXT(3) Manual Page

## Name

vkWriteMicromapsPropertiesEXT - Query micromap meta-data on the host



## <a href="#_c_specification" class="anchor"></a>C Specification

To query micromap size parameters on the host, call:

``` c
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

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device which owns the micromaps in `pMicromaps`.

- `micromapCount` is the count of micromaps for which to query the
  property.

- `pMicromaps` is a pointer to an array of existing previously built
  micromaps.

- `queryType` is a [VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html) value specifying the
  property to be queried.

- `dataSize` is the size in bytes of the buffer pointed to by `pData`.

- `pData` is a pointer to a user-allocated buffer where the results will
  be written.

- `stride` is the stride in bytes between results for individual queries
  within `pData`.

## <a href="#_description" class="anchor"></a>Description

This command fulfills the same task as
[vkCmdWriteMicromapsPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteMicromapsPropertiesEXT.html)
but is executed by the host.

Valid Usage

- <a href="#VUID-vkWriteMicromapsPropertiesEXT-pMicromaps-07501"
  id="VUID-vkWriteMicromapsPropertiesEXT-pMicromaps-07501"></a>
  VUID-vkWriteMicromapsPropertiesEXT-pMicromaps-07501  
  All micromaps in `pMicromaps` **must** have been constructed prior to
  the execution of this command

- <a href="#VUID-vkWriteMicromapsPropertiesEXT-pMicromaps-07502"
  id="VUID-vkWriteMicromapsPropertiesEXT-pMicromaps-07502"></a>
  VUID-vkWriteMicromapsPropertiesEXT-pMicromaps-07502  
  All micromaps in `pMicromaps` **must** have been constructed with
  `VK_BUILD_MICROMAP_ALLOW_COMPACTION_BIT_EXT` if `queryType` is
  `VK_QUERY_TYPE_MICROMAP_COMPACTED_SIZE_EXT`

- <a href="#VUID-vkWriteMicromapsPropertiesEXT-queryType-07503"
  id="VUID-vkWriteMicromapsPropertiesEXT-queryType-07503"></a>
  VUID-vkWriteMicromapsPropertiesEXT-queryType-07503  
  `queryType` **must** be `VK_QUERY_TYPE_MICROMAP_COMPACTED_SIZE_EXT` or
  `VK_QUERY_TYPE_MICROMAP_SERIALIZATION_SIZE_EXT`

- <a href="#VUID-vkWriteMicromapsPropertiesEXT-queryType-07573"
  id="VUID-vkWriteMicromapsPropertiesEXT-queryType-07573"></a>
  VUID-vkWriteMicromapsPropertiesEXT-queryType-07573  
  If `queryType` is `VK_QUERY_TYPE_MICROMAP_SERIALIZATION_SIZE_EXT`,
  then `stride` **must** be a multiple of the size of
  [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)

- <a href="#VUID-vkWriteMicromapsPropertiesEXT-queryType-07574"
  id="VUID-vkWriteMicromapsPropertiesEXT-queryType-07574"></a>
  VUID-vkWriteMicromapsPropertiesEXT-queryType-07574  
  If `queryType` is `VK_QUERY_TYPE_MICROMAP_SERIALIZATION_SIZE_EXT`,
  then `pData` **must** point to a [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)

- <a href="#VUID-vkWriteMicromapsPropertiesEXT-queryType-07575"
  id="VUID-vkWriteMicromapsPropertiesEXT-queryType-07575"></a>
  VUID-vkWriteMicromapsPropertiesEXT-queryType-07575  
  If `queryType` is

- <a href="#VUID-vkWriteMicromapsPropertiesEXT-dataSize-07576"
  id="VUID-vkWriteMicromapsPropertiesEXT-dataSize-07576"></a>
  VUID-vkWriteMicromapsPropertiesEXT-dataSize-07576  
  `dataSize` **must** be greater than or equal to
  `micromapCount`\*`stride`

- <a href="#VUID-vkWriteMicromapsPropertiesEXT-buffer-07577"
  id="VUID-vkWriteMicromapsPropertiesEXT-buffer-07577"></a>
  VUID-vkWriteMicromapsPropertiesEXT-buffer-07577  
  The `buffer` used to create each micromap in `pMicromaps` **must** be
  bound to host-visible device memory

- <a href="#VUID-vkWriteMicromapsPropertiesEXT-micromapHostCommands-07578"
  id="VUID-vkWriteMicromapsPropertiesEXT-micromapHostCommands-07578"></a>
  VUID-vkWriteMicromapsPropertiesEXT-micromapHostCommands-07578  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-micromapHostCommands"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceOpacityMicromapFeaturesEXT</code>::<code>micromapHostCommands</code></a>
  feature **must** be enabled

- <a href="#VUID-vkWriteMicromapsPropertiesEXT-buffer-07579"
  id="VUID-vkWriteMicromapsPropertiesEXT-buffer-07579"></a>
  VUID-vkWriteMicromapsPropertiesEXT-buffer-07579  
  The `buffer` used to create each micromap in `pMicromaps` **must** be
  bound to memory that was not allocated with multiple instances

Valid Usage (Implicit)

- <a href="#VUID-vkWriteMicromapsPropertiesEXT-device-parameter"
  id="VUID-vkWriteMicromapsPropertiesEXT-device-parameter"></a>
  VUID-vkWriteMicromapsPropertiesEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkWriteMicromapsPropertiesEXT-pMicromaps-parameter"
  id="VUID-vkWriteMicromapsPropertiesEXT-pMicromaps-parameter"></a>
  VUID-vkWriteMicromapsPropertiesEXT-pMicromaps-parameter  
  `pMicromaps` **must** be a valid pointer to an array of
  `micromapCount` valid [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapEXT.html) handles

- <a href="#VUID-vkWriteMicromapsPropertiesEXT-queryType-parameter"
  id="VUID-vkWriteMicromapsPropertiesEXT-queryType-parameter"></a>
  VUID-vkWriteMicromapsPropertiesEXT-queryType-parameter  
  `queryType` **must** be a valid [VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html) value

- <a href="#VUID-vkWriteMicromapsPropertiesEXT-pData-parameter"
  id="VUID-vkWriteMicromapsPropertiesEXT-pData-parameter"></a>
  VUID-vkWriteMicromapsPropertiesEXT-pData-parameter  
  `pData` **must** be a valid pointer to an array of `dataSize` bytes

- <a href="#VUID-vkWriteMicromapsPropertiesEXT-micromapCount-arraylength"
  id="VUID-vkWriteMicromapsPropertiesEXT-micromapCount-arraylength"></a>
  VUID-vkWriteMicromapsPropertiesEXT-micromapCount-arraylength  
  `micromapCount` **must** be greater than `0`

- <a href="#VUID-vkWriteMicromapsPropertiesEXT-dataSize-arraylength"
  id="VUID-vkWriteMicromapsPropertiesEXT-dataSize-arraylength"></a>
  VUID-vkWriteMicromapsPropertiesEXT-dataSize-arraylength  
  `dataSize` **must** be greater than `0`

- <a href="#VUID-vkWriteMicromapsPropertiesEXT-pMicromaps-parent"
  id="VUID-vkWriteMicromapsPropertiesEXT-pMicromaps-parent"></a>
  VUID-vkWriteMicromapsPropertiesEXT-pMicromaps-parent  
  Each element of `pMicromaps` **must** have been created, allocated, or
  retrieved from `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapEXT.html),
[VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkWriteMicromapsPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
