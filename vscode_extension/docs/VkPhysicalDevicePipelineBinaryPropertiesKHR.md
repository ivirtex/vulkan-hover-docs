# VkPhysicalDevicePipelineBinaryPropertiesKHR(3) Manual Page

## Name

VkPhysicalDevicePipelineBinaryPropertiesKHR - Structure describing properties about the pipeline binary implementation



## [](#_c_specification)C Specification

The `VkPhysicalDevicePipelineBinaryPropertiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_pipeline_binary
typedef struct VkPhysicalDevicePipelineBinaryPropertiesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           pipelineBinaryInternalCache;
    VkBool32           pipelineBinaryInternalCacheControl;
    VkBool32           pipelineBinaryPrefersInternalCache;
    VkBool32           pipelineBinaryPrecompiledInternalCache;
    VkBool32           pipelineBinaryCompressedData;
} VkPhysicalDevicePipelineBinaryPropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`pipelineBinaryInternalCache` specifies that the implementation maintains a pipeline cache internal to the implementation. If this is `VK_TRUE`, applications **can** create pipeline binaries with only a pipeline create info, and in this case, an implementation **may** be able to create a pipeline binary directly without application needing to capture the binary itself.
- []()`pipelineBinaryInternalCacheControl` specifies whether the driver’s internal cache **can** be disabled. If this property is `VK_TRUE` [VkDevicePipelineBinaryInternalCacheControlKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevicePipelineBinaryInternalCacheControlKHR.html)::`disableInternalCache` **can** be used to disable the driver’s internal cache, allowing an application to take full control of both memory and disk usage.
- []()`pipelineBinaryPrefersInternalCache` specifies that the implementation prefers to maintain an internal cache, and applications **should** not store pipeline binaries in their own on-disk caches to avoid increased on-disk storage requirements. Applications are encouraged to only store pipeline keys instead, and aim to create pipeline binaries from key alone on subsequent runs of the application.
- []()`pipelineBinaryPrecompiledInternalCache` specifies that the implementation **may** have pipeline binaries in its internal cache, which is populated without the application ever having generated that pipeline itself. Applications **can** attempt to create binaries without extracting pipeline binary data from the pipeline prior for a set of pipeline keys, including from previous runs of the application.
- []()`pipelineBinaryCompressedData` specifies that the binary data is already compressed and so applications **should** not attempt to compress it.

## [](#_description)Description

Note

These properties tend to be platform specific and may change depending on external configuration which is outside the scope of this specification. These properties are intended to guide applications when implementations have dedicated caching solutions available. In particular, if the `pipelineBinaryPrefersInternalCache` limit is exposed, relying on the internal cache may provide some advantage compared to an application-specific solution. An application with its own dedicated solution may still use its own caching system even with this limit exposed.

If the `VkPhysicalDevicePipelineBinaryPropertiesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDevicePipelineBinaryPropertiesKHR-sType-sType)VUID-VkPhysicalDevicePipelineBinaryPropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_BINARY_PROPERTIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_pipeline\_binary](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_binary.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDevicePipelineBinaryPropertiesKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0