# VK\_DEFINE\_HANDLE(3) Manual Page

## Name

VK\_DEFINE\_HANDLE - Declare a dispatchable object handle



## [](#_c_specification)C Specification

`VK_DEFINE_HANDLE` defines a [dispatchable handle](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fundamentals-objectmodel-overview) type.

```c++
// Provided by VK_VERSION_1_0
#define VK_DEFINE_HANDLE(object) typedef struct object##_T* object;
```

## [](#_description)Description

- `object` is the name of the resulting C type.

The only dispatchable handle types are those related to device and instance management, such as [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkExternalComputeQueueNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalComputeQueueNV.html), [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_DEFINE_HANDLE)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0