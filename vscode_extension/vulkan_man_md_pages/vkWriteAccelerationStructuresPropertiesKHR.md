# vkWriteAccelerationStructuresPropertiesKHR(3) Manual Page

## Name

vkWriteAccelerationStructuresPropertiesKHR - Query acceleration
structure meta-data on the host



## <a href="#_c_specification" class="anchor"></a>C Specification

To query acceleration structure size parameters on the host, call:

``` c
// Provided by VK_KHR_acceleration_structure
VkResult vkWriteAccelerationStructuresPropertiesKHR(
    VkDevice                                    device,
    uint32_t                                    accelerationStructureCount,
    const VkAccelerationStructureKHR*           pAccelerationStructures,
    VkQueryType                                 queryType,
    size_t                                      dataSize,
    void*                                       pData,
    size_t                                      stride);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device which owns the acceleration structures in
  `pAccelerationStructures`.

- `accelerationStructureCount` is the count of acceleration structures
  for which to query the property.

- `pAccelerationStructures` is a pointer to an array of existing
  previously built acceleration structures.

- `queryType` is a [VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html) value specifying the
  property to be queried.

- `dataSize` is the size in bytes of the buffer pointed to by `pData`.

- `pData` is a pointer to an application-allocated buffer where the
  results will be written.

- `stride` is the stride in bytes between results for individual queries
  within `pData`.

## <a href="#_description" class="anchor"></a>Description

This command fulfills the same task as
[vkCmdWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteAccelerationStructuresPropertiesKHR.html)
but is executed by the host.

Valid Usage

- <a
  href="#VUID-vkWriteAccelerationStructuresPropertiesKHR-accelerationStructureHostCommands-03585"
  id="VUID-vkWriteAccelerationStructuresPropertiesKHR-accelerationStructureHostCommands-03585"></a>
  VUID-vkWriteAccelerationStructuresPropertiesKHR-accelerationStructureHostCommands-03585  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-accelerationStructureHostCommands"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceAccelerationStructureFeaturesKHR</code>::<code>accelerationStructureHostCommands</code></a>
  feature **must** be enabled

<!-- -->

- <a
  href="#VUID-vkWriteAccelerationStructuresPropertiesKHR-pAccelerationStructures-04964"
  id="VUID-vkWriteAccelerationStructuresPropertiesKHR-pAccelerationStructures-04964"></a>
  VUID-vkWriteAccelerationStructuresPropertiesKHR-pAccelerationStructures-04964  
  All acceleration structures in `pAccelerationStructures` **must** have
  been built prior to the execution of this command

- <a
  href="#VUID-vkWriteAccelerationStructuresPropertiesKHR-accelerationStructures-03431"
  id="VUID-vkWriteAccelerationStructuresPropertiesKHR-accelerationStructures-03431"></a>
  VUID-vkWriteAccelerationStructuresPropertiesKHR-accelerationStructures-03431  
  All acceleration structures in `pAccelerationStructures` **must** have
  been built with
  `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_COMPACTION_BIT_KHR` if
  `queryType` is
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_KHR`

- <a
  href="#VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-06742"
  id="VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-06742"></a>
  VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-06742  
  `queryType` **must** be
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SIZE_KHR`,
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_BOTTOM_LEVEL_POINTERS_KHR`,
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_KHR`, or
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_SIZE_KHR`

- <a
  href="#VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-03448"
  id="VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-03448"></a>
  VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-03448  
  If `queryType` is
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_KHR`, then
  `stride` **must** be a multiple of the size of
  [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)

- <a
  href="#VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-03449"
  id="VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-03449"></a>
  VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-03449  
  If `queryType` is
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_KHR`, then
  `pData` **must** point to a [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)

- <a
  href="#VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-03450"
  id="VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-03450"></a>
  VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-03450  
  If `queryType` is
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_SIZE_KHR`, then
  `stride` **must** be a multiple of the size of
  [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)

- <a
  href="#VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-03451"
  id="VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-03451"></a>
  VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-03451  
  If `queryType` is
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_SIZE_KHR`, then
  `pData` **must** point to a [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)

- <a
  href="#VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-06731"
  id="VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-06731"></a>
  VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-06731  
  If `queryType` is `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SIZE_KHR`,
  then `stride` **must** be a multiple of the size of
  [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)

- <a
  href="#VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-06732"
  id="VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-06732"></a>
  VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-06732  
  If `queryType` is `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SIZE_KHR`,
  then `pData` **must** point to a [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)

- <a
  href="#VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-06733"
  id="VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-06733"></a>
  VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-06733  
  If `queryType` is
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_BOTTOM_LEVEL_POINTERS_KHR`,
  then `stride` **must** be a multiple of the size of
  [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)

- <a
  href="#VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-06734"
  id="VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-06734"></a>
  VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-06734  
  If `queryType` is
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_BOTTOM_LEVEL_POINTERS_KHR`,
  then `pData` **must** point to a [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)

- <a
  href="#VUID-vkWriteAccelerationStructuresPropertiesKHR-dataSize-03452"
  id="VUID-vkWriteAccelerationStructuresPropertiesKHR-dataSize-03452"></a>
  VUID-vkWriteAccelerationStructuresPropertiesKHR-dataSize-03452  
  `dataSize` **must** be greater than or equal to
  `accelerationStructureCount`\*`stride`

- <a href="#VUID-vkWriteAccelerationStructuresPropertiesKHR-buffer-03733"
  id="VUID-vkWriteAccelerationStructuresPropertiesKHR-buffer-03733"></a>
  VUID-vkWriteAccelerationStructuresPropertiesKHR-buffer-03733  
  The `buffer` used to create each acceleration structure in
  `pAccelerationStructures` **must** be bound to host-visible device
  memory

- <a href="#VUID-vkWriteAccelerationStructuresPropertiesKHR-buffer-03784"
  id="VUID-vkWriteAccelerationStructuresPropertiesKHR-buffer-03784"></a>
  VUID-vkWriteAccelerationStructuresPropertiesKHR-buffer-03784  
  The `buffer` used to create each acceleration structure in
  `pAccelerationStructures` **must** be bound to memory that was not
  allocated with multiple instances

Valid Usage (Implicit)

- <a
  href="#VUID-vkWriteAccelerationStructuresPropertiesKHR-device-parameter"
  id="VUID-vkWriteAccelerationStructuresPropertiesKHR-device-parameter"></a>
  VUID-vkWriteAccelerationStructuresPropertiesKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkWriteAccelerationStructuresPropertiesKHR-pAccelerationStructures-parameter"
  id="VUID-vkWriteAccelerationStructuresPropertiesKHR-pAccelerationStructures-parameter"></a>
  VUID-vkWriteAccelerationStructuresPropertiesKHR-pAccelerationStructures-parameter  
  `pAccelerationStructures` **must** be a valid pointer to an array of
  `accelerationStructureCount` valid
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html) handles

- <a
  href="#VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-parameter"
  id="VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-parameter"></a>
  VUID-vkWriteAccelerationStructuresPropertiesKHR-queryType-parameter  
  `queryType` **must** be a valid [VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html) value

- <a
  href="#VUID-vkWriteAccelerationStructuresPropertiesKHR-pData-parameter"
  id="VUID-vkWriteAccelerationStructuresPropertiesKHR-pData-parameter"></a>
  VUID-vkWriteAccelerationStructuresPropertiesKHR-pData-parameter  
  `pData` **must** be a valid pointer to an array of `dataSize` bytes

- <a
  href="#VUID-vkWriteAccelerationStructuresPropertiesKHR-accelerationStructureCount-arraylength"
  id="VUID-vkWriteAccelerationStructuresPropertiesKHR-accelerationStructureCount-arraylength"></a>
  VUID-vkWriteAccelerationStructuresPropertiesKHR-accelerationStructureCount-arraylength  
  `accelerationStructureCount` **must** be greater than `0`

- <a
  href="#VUID-vkWriteAccelerationStructuresPropertiesKHR-dataSize-arraylength"
  id="VUID-vkWriteAccelerationStructuresPropertiesKHR-dataSize-arraylength"></a>
  VUID-vkWriteAccelerationStructuresPropertiesKHR-dataSize-arraylength  
  `dataSize` **must** be greater than `0`

- <a
  href="#VUID-vkWriteAccelerationStructuresPropertiesKHR-pAccelerationStructures-parent"
  id="VUID-vkWriteAccelerationStructuresPropertiesKHR-pAccelerationStructures-parent"></a>
  VUID-vkWriteAccelerationStructuresPropertiesKHR-pAccelerationStructures-parent  
  Each element of `pAccelerationStructures` **must** have been created,
  allocated, or retrieved from `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkWriteAccelerationStructuresPropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
