# VkExternalFenceProperties(3) Manual Page

## Name

VkExternalFenceProperties - Structure describing supported external fence handle features



## [](#_c_specification)C Specification

The `VkExternalFenceProperties` structure is defined as:

```c++
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

```c++
// Provided by VK_KHR_external_fence_capabilities
typedef VkExternalFenceProperties VkExternalFencePropertiesKHR;
```

## [](#_members)Members

- `exportFromImportedHandleTypes` is a bitmask of [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceHandleTypeFlagBits.html) indicating which types of imported handle `handleType` **can** be exported from.
- `compatibleHandleTypes` is a bitmask of [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceHandleTypeFlagBits.html) specifying handle types which **can** be specified at the same time as `handleType` when creating a fence.
- `externalFenceFeatures` is a bitmask of [VkExternalFenceFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceFeatureFlagBits.html) indicating the features of `handleType`.

## [](#_description)Description

If `handleType` is not supported by the implementation, then [VkExternalFenceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceProperties.html)::`externalFenceFeatures` will be zero.

Valid Usage (Implicit)

- [](#VUID-VkExternalFenceProperties-sType-sType)VUID-VkExternalFenceProperties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXTERNAL_FENCE_PROPERTIES`
- [](#VUID-VkExternalFenceProperties-pNext-pNext)VUID-VkExternalFenceProperties-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkExternalFenceFeatureFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceFeatureFlags.html), [VkExternalFenceHandleTypeFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceHandleTypeFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPhysicalDeviceExternalFenceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceExternalFenceProperties.html), [vkGetPhysicalDeviceExternalFenceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceExternalFenceProperties.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExternalFenceProperties).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0