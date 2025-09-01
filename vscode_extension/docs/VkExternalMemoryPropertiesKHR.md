# VkExternalMemoryProperties(3) Manual Page

## Name

VkExternalMemoryProperties - Structure specifying external memory handle type capabilities



## [](#_c_specification)C Specification

The `VkExternalMemoryProperties` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkExternalMemoryProperties {
    VkExternalMemoryFeatureFlags       externalMemoryFeatures;
    VkExternalMemoryHandleTypeFlags    exportFromImportedHandleTypes;
    VkExternalMemoryHandleTypeFlags    compatibleHandleTypes;
} VkExternalMemoryProperties;
```

or the equivalent

```c++
// Provided by VK_KHR_external_memory_capabilities
typedef VkExternalMemoryProperties VkExternalMemoryPropertiesKHR;
```

## [](#_members)Members

- `externalMemoryFeatures` is a bitmask of [VkExternalMemoryFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryFeatureFlagBits.html) specifying the features of `handleType`.
- `exportFromImportedHandleTypes` is a bitmask of [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) specifying which types of imported handle `handleType` **can** be exported from.
- `compatibleHandleTypes` is a bitmask of [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) specifying handle types which **can** be specified at the same time as `handleType` when creating an image compatible with external memory.

## [](#_description)Description

`compatibleHandleTypes` **must** include at least `handleType`. Inclusion of a handle type in `compatibleHandleTypes` does not imply the values returned in [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties2.html) will be the same when [VkPhysicalDeviceExternalImageFormatInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalImageFormatInfo.html)::`handleType` is set to that type. The application is responsible for querying the capabilities of all handle types intended for concurrent use in a single image and intersecting them to obtain the compatible set of capabilities.

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkExternalBufferProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalBufferProperties.html), [VkExternalImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalImageFormatProperties.html), [VkExternalMemoryFeatureFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryFeatureFlags.html), [VkExternalMemoryHandleTypeFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlags.html), [VkExternalTensorPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalTensorPropertiesARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExternalMemoryProperties).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0