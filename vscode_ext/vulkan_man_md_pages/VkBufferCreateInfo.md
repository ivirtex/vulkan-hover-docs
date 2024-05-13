# VkBufferCreateInfo(3) Manual Page

## Name

VkBufferCreateInfo - Structure specifying the parameters of a newly
created buffer object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBufferCreateInfo` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkBufferCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateFlagBits.html) specifying
  additional parameters of the buffer.

- `size` is the size in bytes of the buffer to be created.

- `usage` is a bitmask of
  [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits.html) specifying allowed
  usages of the buffer.

- `sharingMode` is a [VkSharingMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSharingMode.html) value
  specifying the sharing mode of the buffer when it will be accessed by
  multiple queue families.

- `queueFamilyIndexCount` is the number of entries in the
  `pQueueFamilyIndices` array.

- `pQueueFamilyIndices` is a pointer to an array of queue families that
  will access this buffer. It is ignored if `sharingMode` is not
  `VK_SHARING_MODE_CONCURRENT`.

## <a href="#_description" class="anchor"></a>Description

If a
[VkBufferUsageFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags2CreateInfoKHR.html)
structure is present in the `pNext` chain,
[VkBufferUsageFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags2CreateInfoKHR.html)::`usage`
from that structure is used instead of `usage` from this structure.

Valid Usage

- <a href="#VUID-VkBufferCreateInfo-None-09499"
  id="VUID-VkBufferCreateInfo-None-09499"></a>
  VUID-VkBufferCreateInfo-None-09499  
  If the `pNext` chain does not include a
  [VkBufferUsageFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags2CreateInfoKHR.html)
  structure, `usage` must be a valid combination of
  [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits.html) values

- <a href="#VUID-VkBufferCreateInfo-None-09500"
  id="VUID-VkBufferCreateInfo-None-09500"></a>
  VUID-VkBufferCreateInfo-None-09500  
  If the `pNext` chain does not include a
  [VkBufferUsageFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags2CreateInfoKHR.html)
  structure, `usage` must not be 0

- <a href="#VUID-VkBufferCreateInfo-size-00912"
  id="VUID-VkBufferCreateInfo-size-00912"></a>
  VUID-VkBufferCreateInfo-size-00912  
  `size` **must** be greater than `0`

- <a href="#VUID-VkBufferCreateInfo-sharingMode-00913"
  id="VUID-VkBufferCreateInfo-sharingMode-00913"></a>
  VUID-VkBufferCreateInfo-sharingMode-00913  
  If `sharingMode` is `VK_SHARING_MODE_CONCURRENT`,
  `pQueueFamilyIndices` **must** be a valid pointer to an array of
  `queueFamilyIndexCount` `uint32_t` values

- <a href="#VUID-VkBufferCreateInfo-sharingMode-00914"
  id="VUID-VkBufferCreateInfo-sharingMode-00914"></a>
  VUID-VkBufferCreateInfo-sharingMode-00914  
  If `sharingMode` is `VK_SHARING_MODE_CONCURRENT`,
  `queueFamilyIndexCount` **must** be greater than `1`

- <a href="#VUID-VkBufferCreateInfo-sharingMode-01419"
  id="VUID-VkBufferCreateInfo-sharingMode-01419"></a>
  VUID-VkBufferCreateInfo-sharingMode-01419  
  If `sharingMode` is `VK_SHARING_MODE_CONCURRENT`, each element of
  `pQueueFamilyIndices` **must** be unique and **must** be less than
  `pQueueFamilyPropertyCount` returned by either
  [vkGetPhysicalDeviceQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html)
  or
  [vkGetPhysicalDeviceQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties.html)
  for the `physicalDevice` that was used to create `device`

- <a href="#VUID-VkBufferCreateInfo-flags-00915"
  id="VUID-VkBufferCreateInfo-flags-00915"></a>
  VUID-VkBufferCreateInfo-flags-00915  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-sparseBinding"
  target="_blank" rel="noopener"><code>sparseBinding</code></a> feature
  is not enabled, `flags` **must** not contain
  `VK_BUFFER_CREATE_SPARSE_BINDING_BIT`

- <a href="#VUID-VkBufferCreateInfo-flags-00916"
  id="VUID-VkBufferCreateInfo-flags-00916"></a>
  VUID-VkBufferCreateInfo-flags-00916  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-sparseResidencyBuffer"
  target="_blank" rel="noopener"><code>sparseResidencyBuffer</code></a>
  feature is not enabled, `flags` **must** not contain
  `VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT`

- <a href="#VUID-VkBufferCreateInfo-flags-00917"
  id="VUID-VkBufferCreateInfo-flags-00917"></a>
  VUID-VkBufferCreateInfo-flags-00917  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-sparseResidencyAliased"
  target="_blank" rel="noopener"><code>sparseResidencyAliased</code></a>
  feature is not enabled, `flags` **must** not contain
  `VK_BUFFER_CREATE_SPARSE_ALIASED_BIT`

- <a href="#VUID-VkBufferCreateInfo-flags-00918"
  id="VUID-VkBufferCreateInfo-flags-00918"></a>
  VUID-VkBufferCreateInfo-flags-00918  
  If `flags` contains `VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT` or
  `VK_BUFFER_CREATE_SPARSE_ALIASED_BIT`, it **must** also contain
  `VK_BUFFER_CREATE_SPARSE_BINDING_BIT`

- <a href="#VUID-VkBufferCreateInfo-pNext-00920"
  id="VUID-VkBufferCreateInfo-pNext-00920"></a>
  VUID-VkBufferCreateInfo-pNext-00920  
  If the `pNext` chain includes a
  [VkExternalMemoryBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryBufferCreateInfo.html)
  structure, its `handleTypes` member **must** only contain bits that
  are also in
  [VkExternalBufferProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalBufferProperties.html)::`externalMemoryProperties.compatibleHandleTypes`,
  as returned by
  [vkGetPhysicalDeviceExternalBufferProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalBufferProperties.html)
  with `pExternalBufferInfo->handleType` equal to any one of the handle
  types specified in
  [VkExternalMemoryBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryBufferCreateInfo.html)::`handleTypes`

- <a href="#VUID-VkBufferCreateInfo-flags-01887"
  id="VUID-VkBufferCreateInfo-flags-01887"></a>
  VUID-VkBufferCreateInfo-flags-01887  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-protectedMemory"
  target="_blank" rel="noopener"><code>protectedMemory</code></a>
  feature is not enabled, `flags` **must** not contain
  `VK_BUFFER_CREATE_PROTECTED_BIT`

- <a href="#VUID-VkBufferCreateInfo-None-01888"
  id="VUID-VkBufferCreateInfo-None-01888"></a>
  VUID-VkBufferCreateInfo-None-01888  
  If any of the bits `VK_BUFFER_CREATE_SPARSE_BINDING_BIT`,
  `VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT`, or
  `VK_BUFFER_CREATE_SPARSE_ALIASED_BIT` are set,
  `VK_BUFFER_CREATE_PROTECTED_BIT` **must** not also be set

- <a href="#VUID-VkBufferCreateInfo-pNext-01571"
  id="VUID-VkBufferCreateInfo-pNext-01571"></a>
  VUID-VkBufferCreateInfo-pNext-01571  
  If the `pNext` chain includes a
  [VkDedicatedAllocationBufferCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDedicatedAllocationBufferCreateInfoNV.html)
  structure, and the `dedicatedAllocation` member of the chained
  structure is `VK_TRUE`, then `flags` **must** not include
  `VK_BUFFER_CREATE_SPARSE_BINDING_BIT`,
  `VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT`, or
  `VK_BUFFER_CREATE_SPARSE_ALIASED_BIT`

- <a href="#VUID-VkBufferCreateInfo-deviceAddress-02604"
  id="VUID-VkBufferCreateInfo-deviceAddress-02604"></a>
  VUID-VkBufferCreateInfo-deviceAddress-02604  
  If
  [VkBufferDeviceAddressCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferDeviceAddressCreateInfoEXT.html)::`deviceAddress`
  is not zero, `flags` **must** include
  `VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT`

- <a href="#VUID-VkBufferCreateInfo-opaqueCaptureAddress-03337"
  id="VUID-VkBufferCreateInfo-opaqueCaptureAddress-03337"></a>
  VUID-VkBufferCreateInfo-opaqueCaptureAddress-03337  
  If
  [VkBufferOpaqueCaptureAddressCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferOpaqueCaptureAddressCreateInfo.html)::`opaqueCaptureAddress`
  is not zero, `flags` **must** include
  `VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT`

- <a href="#VUID-VkBufferCreateInfo-flags-03338"
  id="VUID-VkBufferCreateInfo-flags-03338"></a>
  VUID-VkBufferCreateInfo-flags-03338  
  If `flags` includes
  `VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT`, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-bufferDeviceAddressCaptureReplayEXT"
  target="_blank"
  rel="noopener"><code>bufferDeviceAddressCaptureReplay</code></a>
  feature **must** be enabled

- <a href="#VUID-VkBufferCreateInfo-usage-04813"
  id="VUID-VkBufferCreateInfo-usage-04813"></a>
  VUID-VkBufferCreateInfo-usage-04813  
  If `usage` includes `VK_BUFFER_USAGE_VIDEO_DECODE_SRC_BIT_KHR` or
  `VK_BUFFER_USAGE_VIDEO_DECODE_DST_BIT_KHR`, and `flags` does not
  include `VK_BUFFER_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR`, then the
  `pNext` chain **must** include a
  [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html) structure
  with `profileCount` greater than `0` and `pProfiles` including at
  least one [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html)
  structure with a `videoCodecOperation` member specifying a decode
  operation

- <a href="#VUID-VkBufferCreateInfo-usage-04814"
  id="VUID-VkBufferCreateInfo-usage-04814"></a>
  VUID-VkBufferCreateInfo-usage-04814  
  If `usage` includes `VK_BUFFER_USAGE_VIDEO_ENCODE_SRC_BIT_KHR` or
  `VK_BUFFER_USAGE_VIDEO_ENCODE_DST_BIT_KHR`, and `flags` does not
  include `VK_BUFFER_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR`, then the
  `pNext` chain **must** include a
  [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html) structure
  with `profileCount` greater than `0` and `pProfiles` including at
  least one [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html)
  structure with a `videoCodecOperation` member specifying an encode
  operation

- <a href="#VUID-VkBufferCreateInfo-flags-08325"
  id="VUID-VkBufferCreateInfo-flags-08325"></a>
  VUID-VkBufferCreateInfo-flags-08325  
  If `flags` includes
  `VK_BUFFER_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR`, then <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-videoMaintenance1"
  target="_blank" rel="noopener"><code>videoMaintenance1</code></a>
  **must** be enabled

- <a href="#VUID-VkBufferCreateInfo-size-06409"
  id="VUID-VkBufferCreateInfo-size-06409"></a>
  VUID-VkBufferCreateInfo-size-06409  
  `size` **must** be less than or equal to
  [VkPhysicalDeviceMaintenance4Properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance4Properties.html)::`maxBufferSize`

- <a href="#VUID-VkBufferCreateInfo-usage-08097"
  id="VUID-VkBufferCreateInfo-usage-08097"></a>
  VUID-VkBufferCreateInfo-usage-08097  
  If `usage` includes
  `VK_BUFFER_USAGE_SAMPLER_DESCRIPTOR_BUFFER_BIT_EXT`, creating this
  `VkBuffer` **must** not cause the total required space for all
  currently valid buffers using this flag on the device to exceed
  [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`samplerDescriptorBufferAddressSpaceSize`
  or
  [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`descriptorBufferAddressSpaceSize`

- <a href="#VUID-VkBufferCreateInfo-usage-08098"
  id="VUID-VkBufferCreateInfo-usage-08098"></a>
  VUID-VkBufferCreateInfo-usage-08098  
  If `usage` includes
  `VK_BUFFER_USAGE_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT`, creating this
  `VkBuffer` **must** not cause the total required space for all
  currently valid buffers using this flag on the device to exceed
  [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`resourceDescriptorBufferAddressSpaceSize`
  or
  [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`descriptorBufferAddressSpaceSize`

- <a href="#VUID-VkBufferCreateInfo-flags-08099"
  id="VUID-VkBufferCreateInfo-flags-08099"></a>
  VUID-VkBufferCreateInfo-flags-08099  
  If `flags` includes
  `VK_BUFFER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-descriptorBufferCaptureReplay"
  target="_blank"
  rel="noopener"><code>descriptorBufferCaptureReplay</code></a> feature
  **must** be enabled

- <a href="#VUID-VkBufferCreateInfo-pNext-08100"
  id="VUID-VkBufferCreateInfo-pNext-08100"></a>
  VUID-VkBufferCreateInfo-pNext-08100  
  If the `pNext` chain includes a
  [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html)
  structure, `flags` **must** contain
  `VK_BUFFER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`

- <a href="#VUID-VkBufferCreateInfo-usage-08101"
  id="VUID-VkBufferCreateInfo-usage-08101"></a>
  VUID-VkBufferCreateInfo-usage-08101  
  If `usage` includes
  `VK_BUFFER_USAGE_PUSH_DESCRIPTORS_DESCRIPTOR_BUFFER_BIT_EXT`, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-descriptorBufferPushDescriptors"
  target="_blank"
  rel="noopener"><code>descriptorBufferPushDescriptors</code></a>
  feature **must** be enabled

- <a href="#VUID-VkBufferCreateInfo-usage-08102"
  id="VUID-VkBufferCreateInfo-usage-08102"></a>
  VUID-VkBufferCreateInfo-usage-08102  
  If `usage` includes
  `VK_BUFFER_USAGE_PUSH_DESCRIPTORS_DESCRIPTOR_BUFFER_BIT_EXT` <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-bufferlessPushDescriptors"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceDescriptorBufferPropertiesEXT</code>::<code>bufferlessPushDescriptors</code></a>
  **must** be `VK_FALSE`

- <a href="#VUID-VkBufferCreateInfo-usage-08103"
  id="VUID-VkBufferCreateInfo-usage-08103"></a>
  VUID-VkBufferCreateInfo-usage-08103  
  If `usage` includes
  `VK_BUFFER_USAGE_PUSH_DESCRIPTORS_DESCRIPTOR_BUFFER_BIT_EXT`, `usage`
  **must** contain at least one of
  `VK_BUFFER_USAGE_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT` or
  `VK_BUFFER_USAGE_SAMPLER_DESCRIPTOR_BUFFER_BIT_EXT`

- <a href="#VUID-VkBufferCreateInfo-flags-09641"
  id="VUID-VkBufferCreateInfo-flags-09641"></a>
  VUID-VkBufferCreateInfo-flags-09641  
  If `flags` includes `VK_BUFFER_CREATE_PROTECTED_BIT`, then `usage`
  **must** not contain any of the following bits

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

- <a href="#VUID-VkBufferCreateInfo-sType-sType"
  id="VUID-VkBufferCreateInfo-sType-sType"></a>
  VUID-VkBufferCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO`

- <a href="#VUID-VkBufferCreateInfo-pNext-pNext"
  id="VUID-VkBufferCreateInfo-pNext-pNext"></a>
  VUID-VkBufferCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkBufferCollectionBufferCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionBufferCreateInfoFUCHSIA.html),
  [VkBufferDeviceAddressCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferDeviceAddressCreateInfoEXT.html),
  [VkBufferOpaqueCaptureAddressCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferOpaqueCaptureAddressCreateInfo.html),
  [VkBufferUsageFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags2CreateInfoKHR.html),
  [VkDedicatedAllocationBufferCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDedicatedAllocationBufferCreateInfoNV.html),
  [VkExternalMemoryBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryBufferCreateInfo.html),
  [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html),
  or [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html)

- <a href="#VUID-VkBufferCreateInfo-sType-unique"
  id="VUID-VkBufferCreateInfo-sType-unique"></a>
  VUID-VkBufferCreateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkBufferCreateInfo-flags-parameter"
  id="VUID-VkBufferCreateInfo-flags-parameter"></a>
  VUID-VkBufferCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of
  [VkBufferCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateFlagBits.html) values

- <a href="#VUID-VkBufferCreateInfo-sharingMode-parameter"
  id="VUID-VkBufferCreateInfo-sharingMode-parameter"></a>
  VUID-VkBufferCreateInfo-sharingMode-parameter  
  `sharingMode` **must** be a valid [VkSharingMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSharingMode.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkBufferConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferConstraintsInfoFUCHSIA.html),
[VkBufferCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateFlags.html),
[VkBufferUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags.html),
[VkDeviceBufferMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceBufferMemoryRequirements.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html), [VkSharingMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSharingMode.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBufferCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
