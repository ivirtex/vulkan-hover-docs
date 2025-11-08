# VkBufferCreateInfo(3) Manual Page

## Name

VkBufferCreateInfo - Structure specifying the parameters of a newly created buffer object



## [](#_c_specification)C Specification

The `VkBufferCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkBufferCreateInfo {
    VkStructureType        sType;
    const void*            pNext;
    VkBufferCreateFlags    flags;
    VkDeviceSize           size;
    VkBufferUsageFlags     usage;
    VkSharingMode          sharingMode;
    uint32_t               queueFamilyIndexCount;
    const uint32_t*        pQueueFamilyIndices;
} VkBufferCreateInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkBufferCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateFlagBits.html) specifying additional parameters of the buffer.
- `size` is the size in bytes of the buffer to be created.
- `usage` is a bitmask of [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits.html) specifying allowed usages of the buffer.
- `sharingMode` is a [VkSharingMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSharingMode.html) value specifying the sharing mode of the buffer when it will be accessed by multiple queue families.
- `queueFamilyIndexCount` is the number of entries in the `pQueueFamilyIndices` array.
- `pQueueFamilyIndices` is a pointer to an array of queue families that will access this buffer. It is ignored if `sharingMode` is not `VK_SHARING_MODE_CONCURRENT`.

## [](#_description)Description

If the `pNext` chain includes a [VkBufferUsageFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2CreateInfo.html) structure, [VkBufferUsageFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2CreateInfo.html)::`usage` from that structure is used instead of `usage` from this structure.

Valid Usage

- [](#VUID-VkBufferCreateInfo-None-09499)VUID-VkBufferCreateInfo-None-09499  
  If the `pNext` chain does not include a [VkBufferUsageFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2CreateInfo.html) structure, `usage` **must** be a valid combination of [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits.html) values
- [](#VUID-VkBufferCreateInfo-None-09500)VUID-VkBufferCreateInfo-None-09500  
  If the `pNext` chain does not include a [VkBufferUsageFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2CreateInfo.html) structure, `usage` **must** not be 0
- [](#VUID-VkBufferCreateInfo-size-00912)VUID-VkBufferCreateInfo-size-00912  
  `size` **must** be greater than `0`
- [](#VUID-VkBufferCreateInfo-sharingMode-00913)VUID-VkBufferCreateInfo-sharingMode-00913  
  If `sharingMode` is `VK_SHARING_MODE_CONCURRENT`, `pQueueFamilyIndices` **must** be a valid pointer to an array of `queueFamilyIndexCount` `uint32_t` values
- [](#VUID-VkBufferCreateInfo-sharingMode-00914)VUID-VkBufferCreateInfo-sharingMode-00914  
  If `sharingMode` is `VK_SHARING_MODE_CONCURRENT`, `queueFamilyIndexCount` **must** be greater than `1`
- [](#VUID-VkBufferCreateInfo-sharingMode-01419)VUID-VkBufferCreateInfo-sharingMode-01419  
  If `sharingMode` is `VK_SHARING_MODE_CONCURRENT`, each element of `pQueueFamilyIndices` **must** be unique and **must** be less than `pQueueFamilyPropertyCount` returned by either [vkGetPhysicalDeviceQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html) or [vkGetPhysicalDeviceQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyProperties.html) for the `physicalDevice` that was used to create `device`
- [](#VUID-VkBufferCreateInfo-flags-00915)VUID-VkBufferCreateInfo-flags-00915  
  If the [`sparseBinding`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-sparseBinding) feature is not enabled, `flags` **must** not contain `VK_BUFFER_CREATE_SPARSE_BINDING_BIT`
- [](#VUID-VkBufferCreateInfo-flags-00916)VUID-VkBufferCreateInfo-flags-00916  
  If the [`sparseResidencyBuffer`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-sparseResidencyBuffer) feature is not enabled, `flags` **must** not contain `VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT`
- [](#VUID-VkBufferCreateInfo-flags-00917)VUID-VkBufferCreateInfo-flags-00917  
  If the [`sparseResidencyAliased`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-sparseResidencyAliased) feature is not enabled, `flags` **must** not contain `VK_BUFFER_CREATE_SPARSE_ALIASED_BIT`
- [](#VUID-VkBufferCreateInfo-flags-00918)VUID-VkBufferCreateInfo-flags-00918  
  If `flags` contains `VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT` or `VK_BUFFER_CREATE_SPARSE_ALIASED_BIT`, it **must** also contain `VK_BUFFER_CREATE_SPARSE_BINDING_BIT`
- [](#VUID-VkBufferCreateInfo-pNext-00920)VUID-VkBufferCreateInfo-pNext-00920  
  If the `pNext` chain includes a [VkExternalMemoryBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryBufferCreateInfo.html) structure, its `handleTypes` member **must** only contain bits that are also in [VkExternalBufferProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalBufferProperties.html)::`externalMemoryProperties.compatibleHandleTypes`, as returned by [vkGetPhysicalDeviceExternalBufferProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceExternalBufferProperties.html) with `pExternalBufferInfo->handleType` equal to any one of the handle types specified in [VkExternalMemoryBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryBufferCreateInfo.html)::`handleTypes`
- [](#VUID-VkBufferCreateInfo-flags-01887)VUID-VkBufferCreateInfo-flags-01887  
  If the [`protectedMemory`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-protectedMemory) feature is not enabled, `flags` **must** not contain `VK_BUFFER_CREATE_PROTECTED_BIT`
- [](#VUID-VkBufferCreateInfo-None-01888)VUID-VkBufferCreateInfo-None-01888  
  If any of the bits `VK_BUFFER_CREATE_SPARSE_BINDING_BIT`, `VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT`, or `VK_BUFFER_CREATE_SPARSE_ALIASED_BIT` are set, `VK_BUFFER_CREATE_PROTECTED_BIT` **must** not also be set
- [](#VUID-VkBufferCreateInfo-pNext-01571)VUID-VkBufferCreateInfo-pNext-01571  
  If the `pNext` chain includes a [VkDedicatedAllocationBufferCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDedicatedAllocationBufferCreateInfoNV.html) structure, and the `dedicatedAllocation` member of the chained structure is `VK_TRUE`, then `flags` **must** not include `VK_BUFFER_CREATE_SPARSE_BINDING_BIT`, `VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT`, or `VK_BUFFER_CREATE_SPARSE_ALIASED_BIT`
- [](#VUID-VkBufferCreateInfo-deviceAddress-02604)VUID-VkBufferCreateInfo-deviceAddress-02604  
  If [VkBufferDeviceAddressCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferDeviceAddressCreateInfoEXT.html)::`deviceAddress` is not zero, `flags` **must** include `VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT`
- [](#VUID-VkBufferCreateInfo-opaqueCaptureAddress-03337)VUID-VkBufferCreateInfo-opaqueCaptureAddress-03337  
  If [VkBufferOpaqueCaptureAddressCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferOpaqueCaptureAddressCreateInfo.html)::`opaqueCaptureAddress` is not zero, `flags` **must** include `VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT`
- [](#VUID-VkBufferCreateInfo-flags-03338)VUID-VkBufferCreateInfo-flags-03338  
  If `flags` includes `VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT`, the [VkPhysicalDeviceBufferDeviceAddressFeaturesEXT:`bufferDeviceAddressCaptureReplay`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddressCaptureReplayEXT) feature or the [`bufferDeviceAddressCaptureReplay`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddressCaptureReplay) feature **must** be enabled
- [](#VUID-VkBufferCreateInfo-usage-04813)VUID-VkBufferCreateInfo-usage-04813  
  If `usage` includes `VK_BUFFER_USAGE_VIDEO_DECODE_SRC_BIT_KHR` or `VK_BUFFER_USAGE_VIDEO_DECODE_DST_BIT_KHR`, and `flags` does not include `VK_BUFFER_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR`, then the `pNext` chain **must** include a [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileListInfoKHR.html) structure with `profileCount` greater than `0` and `pProfiles` including at least one [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html) structure with a `videoCodecOperation` member specifying a decode operation
- [](#VUID-VkBufferCreateInfo-usage-04814)VUID-VkBufferCreateInfo-usage-04814  
  If `usage` includes `VK_BUFFER_USAGE_VIDEO_ENCODE_SRC_BIT_KHR` or `VK_BUFFER_USAGE_VIDEO_ENCODE_DST_BIT_KHR`, and `flags` does not include `VK_BUFFER_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR`, then the `pNext` chain **must** include a [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileListInfoKHR.html) structure with `profileCount` greater than `0` and `pProfiles` including at least one [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html) structure with a `videoCodecOperation` member specifying an encode operation
- [](#VUID-VkBufferCreateInfo-flags-08325)VUID-VkBufferCreateInfo-flags-08325  
  If `flags` includes `VK_BUFFER_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR`, then [`videoMaintenance1`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-videoMaintenance1) **must** be enabled
- [](#VUID-VkBufferCreateInfo-pNext-10783)VUID-VkBufferCreateInfo-pNext-10783  
  If the `pNext` chain includes a [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileListInfoKHR.html) structure and for any element of its `pProfiles` member `videoCodecOperation` is `VK_VIDEO_CODEC_OPERATION_DECODE_VP9_BIT_KHR`, then the [`videoDecodeVP9`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-videoDecodeVP9) feature **must** be enabled
- [](#VUID-VkBufferCreateInfo-pNext-10249)VUID-VkBufferCreateInfo-pNext-10249  
  If the `pNext` chain includes a [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileListInfoKHR.html) structure and for any element of its `pProfiles` member `videoCodecOperation` is `VK_VIDEO_CODEC_OPERATION_ENCODE_AV1_BIT_KHR`, then the [`videoEncodeAV1`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-videoEncodeAV1) feature **must** be enabled
- [](#VUID-VkBufferCreateInfo-pNext-10919)VUID-VkBufferCreateInfo-pNext-10919  
  If the `pNext` chain includes a [VkVideoEncodeProfileRgbConversionInfoVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeProfileRgbConversionInfoVALVE.html) structure, then the [`videoEncodeRgbConversion`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-videoEncodeRgbConversion) feature **must** be enabled
- [](#VUID-VkBufferCreateInfo-size-06409)VUID-VkBufferCreateInfo-size-06409  
  `size` **must** be less than or equal to [VkPhysicalDeviceMaintenance4Properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance4Properties.html)::`maxBufferSize`
- [](#VUID-VkBufferCreateInfo-usage-08097)VUID-VkBufferCreateInfo-usage-08097  
  If `usage` includes `VK_BUFFER_USAGE_SAMPLER_DESCRIPTOR_BUFFER_BIT_EXT`, creating this `VkBuffer` **must** not cause the total required space for all currently valid buffers using this flag on the device to exceed [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`samplerDescriptorBufferAddressSpaceSize` or [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`descriptorBufferAddressSpaceSize`
- [](#VUID-VkBufferCreateInfo-usage-08098)VUID-VkBufferCreateInfo-usage-08098  
  If `usage` includes `VK_BUFFER_USAGE_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT`, creating this `VkBuffer` **must** not cause the total required space for all currently valid buffers using this flag on the device to exceed [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`resourceDescriptorBufferAddressSpaceSize` or [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`descriptorBufferAddressSpaceSize`
- [](#VUID-VkBufferCreateInfo-flags-08099)VUID-VkBufferCreateInfo-flags-08099  
  If `flags` includes `VK_BUFFER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`, the [`descriptorBufferCaptureReplay`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBufferCaptureReplay) feature **must** be enabled
- [](#VUID-VkBufferCreateInfo-pNext-08100)VUID-VkBufferCreateInfo-pNext-08100  
  If the `pNext` chain includes a [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html) structure, `flags` **must** contain `VK_BUFFER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`
- [](#VUID-VkBufferCreateInfo-usage-08101)VUID-VkBufferCreateInfo-usage-08101  
  If `usage` includes `VK_BUFFER_USAGE_PUSH_DESCRIPTORS_DESCRIPTOR_BUFFER_BIT_EXT`, the [`descriptorBufferPushDescriptors`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBufferPushDescriptors) feature **must** be enabled
- [](#VUID-VkBufferCreateInfo-usage-08102)VUID-VkBufferCreateInfo-usage-08102  
  If `usage` includes `VK_BUFFER_USAGE_PUSH_DESCRIPTORS_DESCRIPTOR_BUFFER_BIT_EXT` [`VkPhysicalDeviceDescriptorBufferPropertiesEXT`::`bufferlessPushDescriptors`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-bufferlessPushDescriptors) **must** be `VK_FALSE`
- [](#VUID-VkBufferCreateInfo-usage-08103)VUID-VkBufferCreateInfo-usage-08103  
  If `usage` includes `VK_BUFFER_USAGE_PUSH_DESCRIPTORS_DESCRIPTOR_BUFFER_BIT_EXT`, `usage` **must** contain at least one of `VK_BUFFER_USAGE_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT` or `VK_BUFFER_USAGE_SAMPLER_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-VkBufferCreateInfo-tileMemoryHeap-10762)VUID-VkBufferCreateInfo-tileMemoryHeap-10762  
  If the [`tileMemoryHeap`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-tileMemoryHeap) feature is not enabled, `usage` **must** not include `VK_BUFFER_USAGE_TILE_MEMORY_BIT_QCOM`
- [](#VUID-VkBufferCreateInfo-usage-10763)VUID-VkBufferCreateInfo-usage-10763  
  If `usage` includes `VK_BUFFER_USAGE_TILE_MEMORY_BIT_QCOM`, then `flags` **must** not contain any of the following bits
  
  - `VK_BUFFER_CREATE_SPARSE_BINDING_BIT`
  - `VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT`
  - `VK_BUFFER_CREATE_SPARSE_ALIASED_BIT`
  - `VK_BUFFER_CREATE_PROTECTED_BIT`
  - `VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT`
  - `VK_BUFFER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`
  - `VK_BUFFER_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR`
- [](#VUID-VkBufferCreateInfo-usage-10764)VUID-VkBufferCreateInfo-usage-10764  
  If `usage` includes `VK_BUFFER_USAGE_TILE_MEMORY_BIT_QCOM`, then only the following `usages` may be set:
  
  - `VK_BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT`
  - `VK_BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT`
  - `VK_BUFFER_USAGE_UNIFORM_BUFFER_BIT`
  - `VK_BUFFER_USAGE_STORAGE_BUFFER_BIT`
  - `VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT`
  - and if [VkPhysicalDeviceTileMemoryHeapPropertiesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTileMemoryHeapPropertiesQCOM.html)::`tileBufferTransfers` is `VK_TRUE` then additionally `VK_BUFFER_USAGE_TRANSFER_SRC_BIT` or `VK_BUFFER_USAGE_TRANSFER_DST_BIT`
- [](#VUID-VkBufferCreateInfo-flags-09641)VUID-VkBufferCreateInfo-flags-09641  
  If `flags` includes `VK_BUFFER_CREATE_PROTECTED_BIT`, then `usage` **must** not contain any of the following bits
  
  - `VK_BUFFER_USAGE_TRANSFORM_FEEDBACK_BUFFER_BIT_EXT`
  - `VK_BUFFER_USAGE_TRANSFORM_FEEDBACK_COUNTER_BUFFER_BIT_EXT`
  - `VK_BUFFER_USAGE_CONDITIONAL_RENDERING_BIT_EXT`
  - `VK_BUFFER_USAGE_ACCELERATION_STRUCTURE_BUILD_INPUT_READ_ONLY_BIT_KHR`
  - `VK_BUFFER_USAGE_ACCELERATION_STRUCTURE_STORAGE_BIT_KHR`
  - `VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR`
  - `VK_BUFFER_USAGE_SAMPLER_DESCRIPTOR_BUFFER_BIT_EXT`
  - `VK_BUFFER_USAGE_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT`
  - `VK_BUFFER_USAGE_PUSH_DESCRIPTORS_DESCRIPTOR_BUFFER_BIT_EXT`
  - `VK_BUFFER_USAGE_MICROMAP_BUILD_INPUT_READ_ONLY_BIT_EXT`
  - `VK_BUFFER_USAGE_MICROMAP_STORAGE_BIT_EXT`

Valid Usage (Implicit)

- [](#VUID-VkBufferCreateInfo-sType-sType)VUID-VkBufferCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO`
- [](#VUID-VkBufferCreateInfo-pNext-pNext)VUID-VkBufferCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkBufferCollectionBufferCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionBufferCreateInfoFUCHSIA.html), [VkBufferDeviceAddressCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferDeviceAddressCreateInfoEXT.html), [VkBufferOpaqueCaptureAddressCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferOpaqueCaptureAddressCreateInfo.html), [VkBufferUsageFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2CreateInfo.html), [VkDedicatedAllocationBufferCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDedicatedAllocationBufferCreateInfoNV.html), [VkExternalMemoryBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryBufferCreateInfo.html), [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html), or [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileListInfoKHR.html)
- [](#VUID-VkBufferCreateInfo-sType-unique)VUID-VkBufferCreateInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkBufferCreateInfo-flags-parameter)VUID-VkBufferCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkBufferCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateFlagBits.html) values
- [](#VUID-VkBufferCreateInfo-sharingMode-parameter)VUID-VkBufferCreateInfo-sharingMode-parameter  
  `sharingMode` **must** be a valid [VkSharingMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSharingMode.html) value

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBufferConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferConstraintsInfoFUCHSIA.html), [VkBufferCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateFlags.html), [VkBufferUsageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags.html), [VkDeviceBufferMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceBufferMemoryRequirements.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkSharingMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSharingMode.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBufferCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0