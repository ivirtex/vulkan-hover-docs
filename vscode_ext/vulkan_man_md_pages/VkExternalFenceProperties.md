# VkExternalFenceProperties(3) Manual Page

## Name

VkExternalFenceProperties - Structure describing supported external
fence handle features



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkExternalFenceProperties` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkExternalFenceProperties {
    VkStructureType                   sType;
    void*                             pNext;
    VkExternalFenceHandleTypeFlags    exportFromImportedHandleTypes;
    VkExternalFenceHandleTypeFlags    compatibleHandleTypes;
    VkExternalFenceFeatureFlags       externalFenceFeatures;
} VkExternalFenceProperties;
```

or the equivalent

``` c
// Provided by VK_KHR_external_fence_capabilities
typedef VkExternalFenceProperties VkExternalFencePropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `exportFromImportedHandleTypes` is a bitmask of
  [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html)
  indicating which types of imported handle `handleType` **can** be
  exported from.

- `compatibleHandleTypes` is a bitmask of
  [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html)
  specifying handle types which **can** be specified at the same time as
  `handleType` when creating a fence.

- `externalFenceFeatures` is a bitmask of
  [VkExternalFenceFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceFeatureFlagBits.html)
  indicating the features of `handleType`.

## <a href="#_description" class="anchor"></a>Description

If `handleType` is not supported by the implementation, then
[VkExternalFenceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceProperties.html)::`externalFenceFeatures`
will be set to zero.

Valid Usage (Implicit)

- <a href="#VUID-VkExternalFenceProperties-sType-sType"
  id="VUID-VkExternalFenceProperties-sType-sType"></a>
  VUID-VkExternalFenceProperties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXTERNAL_FENCE_PROPERTIES`

- <a href="#VUID-VkExternalFenceProperties-pNext-pNext"
  id="VUID-VkExternalFenceProperties-pNext-pNext"></a>
  VUID-VkExternalFenceProperties-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkExternalFenceFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceFeatureFlags.html),
[VkExternalFenceHandleTypeFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPhysicalDeviceExternalFenceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalFenceProperties.html),
[vkGetPhysicalDeviceExternalFencePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalFencePropertiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExternalFenceProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
