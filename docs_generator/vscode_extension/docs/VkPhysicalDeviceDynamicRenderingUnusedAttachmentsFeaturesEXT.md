# VkPhysicalDeviceDynamicRenderingUnusedAttachmentsFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceDynamicRenderingUnusedAttachmentsFeaturesEXT - Structure describing the dynamic rendering unused attachment features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceDynamicRenderingUnusedAttachmentsFeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_dynamic_rendering_unused_attachments
typedef struct VkPhysicalDeviceDynamicRenderingUnusedAttachmentsFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           dynamicRenderingUnusedAttachments;
} VkPhysicalDeviceDynamicRenderingUnusedAttachmentsFeaturesEXT;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`dynamicRenderingUnusedAttachments` indicates that the implementation supports binding graphics pipelines within a render pass instance where any pipeline [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRenderingCreateInfo.html)::`pColorAttachmentFormats` element with a format other than `VK_FORMAT_UNDEFINED` is allowed with a corresponding [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html)::`pColorAttachments` element with an `imageView` equal to [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), or any pipeline [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRenderingCreateInfo.html)::`pColorAttachmentFormats` element with a `VK_FORMAT_UNDEFINED` format is allowed with a corresponding [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html)::`pColorAttachments` element with a non-[VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) `imageView`. Also a [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRenderingCreateInfo.html)::`depthAttachmentFormat` other than `VK_FORMAT_UNDEFINED` is allowed with a [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html)::`pDepthAttachment`, or a [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRenderingCreateInfo.html)::`depthAttachmentFormat` of `VK_FORMAT_UNDEFINED` is allowed with a non-[VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html)::`pDepthAttachment`. Also a [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRenderingCreateInfo.html)::`stencilAttachmentFormat` other than `VK_FORMAT_UNDEFINED` is allowed with a [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html)::`pStencilAttachment`, or a [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRenderingCreateInfo.html)::`stencilAttachmentFormat` of `VK_FORMAT_UNDEFINED` is allowed with a non-[VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html)::`pStencilAttachment`. Any writes to a [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html)::`pColorAttachments`, [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html)::`pDepthAttachment`, or [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html)::`pStencilAttachment` with [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) are discarded.

## [](#_description)Description

If the `VkPhysicalDeviceDynamicRenderingUnusedAttachmentsFeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceDynamicRenderingUnusedAttachmentsFeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceDynamicRenderingUnusedAttachmentsFeaturesEXT-sType-sType)VUID-VkPhysicalDeviceDynamicRenderingUnusedAttachmentsFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DYNAMIC_RENDERING_UNUSED_ATTACHMENTS_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_dynamic\_rendering\_unused\_attachments](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_dynamic_rendering_unused_attachments.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceDynamicRenderingUnusedAttachmentsFeaturesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0