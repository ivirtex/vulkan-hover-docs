# VkExternalSemaphoreProperties(3) Manual Page

## Name

VkExternalSemaphoreProperties - Structure describing supported external
semaphore handle features



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkExternalSemaphoreProperties` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkExternalSemaphoreProperties {
    VkStructureType                       sType;
    void*                                 pNext;
    VkExternalSemaphoreHandleTypeFlags    exportFromImportedHandleTypes;
    VkExternalSemaphoreHandleTypeFlags    compatibleHandleTypes;
    VkExternalSemaphoreFeatureFlags       externalSemaphoreFeatures;
} VkExternalSemaphoreProperties;
```

or the equivalent

``` c
// Provided by VK_KHR_external_semaphore_capabilities
typedef VkExternalSemaphoreProperties VkExternalSemaphorePropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `exportFromImportedHandleTypes` is a bitmask of
  [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html)
  specifying which types of imported handle `handleType` **can** be
  exported from.

- `compatibleHandleTypes` is a bitmask of
  [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html)
  specifying handle types which **can** be specified at the same time as
  `handleType` when creating a semaphore.

- `externalSemaphoreFeatures` is a bitmask of
  [VkExternalSemaphoreFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreFeatureFlagBits.html)
  describing the features of `handleType`.

## <a href="#_description" class="anchor"></a>Description

If `handleType` is not supported by the implementation, then
[VkExternalSemaphoreProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreProperties.html)::`externalSemaphoreFeatures`
will be set to zero.

Valid Usage (Implicit)

- <a href="#VUID-VkExternalSemaphoreProperties-sType-sType"
  id="VUID-VkExternalSemaphoreProperties-sType-sType"></a>
  VUID-VkExternalSemaphoreProperties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES`

- <a href="#VUID-VkExternalSemaphoreProperties-pNext-pNext"
  id="VUID-VkExternalSemaphoreProperties-pNext-pNext"></a>
  VUID-VkExternalSemaphoreProperties-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkExternalSemaphoreFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreFeatureFlags.html),
[VkExternalSemaphoreHandleTypeFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPhysicalDeviceExternalSemaphoreProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalSemaphoreProperties.html),
[vkGetPhysicalDeviceExternalSemaphorePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalSemaphorePropertiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExternalSemaphoreProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
