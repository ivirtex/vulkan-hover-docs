# VkSubpassDescriptionDepthStencilResolve(3) Manual Page

## Name

VkSubpassDescriptionDepthStencilResolve - Structure specifying depth/stencil resolve operations for a subpass



## [](#_c_specification)C Specification

The `VkSubpassDescriptionDepthStencilResolve` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_VERSION_1_2
typedef struct VkSubpassDescriptionDepthStencilResolve {
    VkStructureType                  sType;
    const void*                      pNext;
    VkResolveModeFlagBits            depthResolveMode;
    VkResolveModeFlagBits            stencilResolveMode;
    const VkAttachmentReference2*    pDepthStencilResolveAttachment;
} VkSubpassDescriptionDepthStencilResolve;
```

or the equivalent

```c++
// Provided by VK_KHR_depth_stencil_resolve
typedef VkSubpassDescriptionDepthStencilResolve VkSubpassDescriptionDepthStencilResolveKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `depthResolveMode` is a [VkResolveModeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResolveModeFlagBits.html) value describing the depth resolve mode.
- `stencilResolveMode` is a [VkResolveModeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResolveModeFlagBits.html) value describing the stencil resolve mode.
- `pDepthStencilResolveAttachment` is `NULL` or a pointer to a [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference2.html) structure defining the depth/stencil resolve attachment for this subpass and its layout.

## [](#_description)Description

If the `pNext` chain of [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription2.html) includes a `VkSubpassDescriptionDepthStencilResolve` structure, then that structure describes [multisample resolve operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-resolve-operations) for the depth/stencil attachment in a subpass. If this structure is not included in the `pNext` chain of [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription2.html), or if it is and either `pDepthStencilResolveAttachment` is `NULL` or its attachment index is `VK_ATTACHMENT_UNUSED`, it indicates that no depth/stencil resolve attachment will be used in the subpass.

Valid Usage

- [](#VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03177)VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03177  
  If `pDepthStencilResolveAttachment` is not `NULL` and does not have the value `VK_ATTACHMENT_UNUSED`, `pDepthStencilAttachment` **must** not be `NULL` or have the value `VK_ATTACHMENT_UNUSED`
- [](#VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03179)VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03179  
  If `pDepthStencilResolveAttachment` is not `NULL` and does not have the value `VK_ATTACHMENT_UNUSED`, `pDepthStencilAttachment` **must** not have a sample count of `VK_SAMPLE_COUNT_1_BIT`
- [](#VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03180)VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03180  
  If `pDepthStencilResolveAttachment` is not `NULL` and does not have the value `VK_ATTACHMENT_UNUSED`, `pDepthStencilResolveAttachment` **must** have a sample count of `VK_SAMPLE_COUNT_1_BIT`
- [](#VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-02651)VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-02651  
  If `pDepthStencilResolveAttachment` is not `NULL` and does not have the value `VK_ATTACHMENT_UNUSED` then it **must** have an image format whose [potential format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#potential-format-features) contain `VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03181)VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03181  
  If `pDepthStencilResolveAttachment` is not `NULL` and does not have the value `VK_ATTACHMENT_UNUSED` and [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `pDepthStencilResolveAttachment` has a depth component, then the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `pDepthStencilAttachment` **must** have a depth component with the same number of bits and [numeric format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-numericformat)
- [](#VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03182)VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03182  
  If `pDepthStencilResolveAttachment` is not `NULL` and does not have the value `VK_ATTACHMENT_UNUSED`, and [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `pDepthStencilResolveAttachment` has a stencil component, then the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `pDepthStencilAttachment` **must** have a stencil component with the same number of bits and [numeric format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-numericformat)
- [](#VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03178)VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03178  
  If `pDepthStencilResolveAttachment` is not `NULL` and does not have the value `VK_ATTACHMENT_UNUSED`, `depthResolveMode` and `stencilResolveMode` **must** not both be `VK_RESOLVE_MODE_NONE`
- [](#VUID-VkSubpassDescriptionDepthStencilResolve-depthResolveMode-03183)VUID-VkSubpassDescriptionDepthStencilResolve-depthResolveMode-03183  
  If `pDepthStencilResolveAttachment` is not `NULL` and does not have the value `VK_ATTACHMENT_UNUSED` and the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `pDepthStencilResolveAttachment` has a depth component, then the value of `depthResolveMode` **must** be one of the bits set in [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`supportedDepthResolveModes` or `VK_RESOLVE_MODE_NONE`
- [](#VUID-VkSubpassDescriptionDepthStencilResolve-stencilResolveMode-03184)VUID-VkSubpassDescriptionDepthStencilResolve-stencilResolveMode-03184  
  If `pDepthStencilResolveAttachment` is not `NULL` and does not have the value `VK_ATTACHMENT_UNUSED` and the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `pDepthStencilResolveAttachment` has a stencil component, then the value of `stencilResolveMode` **must** be one of the bits set in [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`supportedStencilResolveModes` or `VK_RESOLVE_MODE_NONE`
- [](#VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03185)VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03185  
  If `pDepthStencilResolveAttachment` is not `NULL` and does not have the value `VK_ATTACHMENT_UNUSED`, the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `pDepthStencilResolveAttachment` has both depth and stencil components, [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`independentResolve` is `VK_FALSE`, and [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`independentResolveNone` is `VK_FALSE`, then the values of `depthResolveMode` and `stencilResolveMode` **must** be identical
- [](#VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03186)VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-03186  
  If `pDepthStencilResolveAttachment` is not `NULL` and does not have the value `VK_ATTACHMENT_UNUSED`, the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `pDepthStencilResolveAttachment` has both depth and stencil components, [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`independentResolve` is `VK_FALSE` and [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`independentResolveNone` is `VK_TRUE`, then the values of `depthResolveMode` and `stencilResolveMode` **must** be identical or one of them **must** be `VK_RESOLVE_MODE_NONE`
- [](#VUID-VkSubpassDescriptionDepthStencilResolve-pNext-06873)VUID-VkSubpassDescriptionDepthStencilResolve-pNext-06873  
  If the `pNext` chain of [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription2.html) includes a `VkMultisampledRenderToSingleSampledInfoEXT` structure, the `multisampledRenderToSingleSampledEnable` field is `VK_TRUE`, and `pDepthStencilAttachment` is not `NULL` and does not have the value `VK_ATTACHMENT_UNUSED`, `depthResolveMode` and `stencilResolveMode` **must** not both be `VK_RESOLVE_MODE_NONE`
- [](#VUID-VkSubpassDescriptionDepthStencilResolve-pNext-06874)VUID-VkSubpassDescriptionDepthStencilResolve-pNext-06874  
  If the `pNext` chain of [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription2.html) includes a `VkMultisampledRenderToSingleSampledInfoEXT` structure whose `multisampledRenderToSingleSampledEnable` field is `VK_TRUE`, and `pDepthStencilAttachment` is not `NULL`, does not have the value `VK_ATTACHMENT_UNUSED`, and has a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) that has a depth component, then the value of `depthResolveMode` **must** be one of the bits set in [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`supportedDepthResolveModes` or `VK_RESOLVE_MODE_NONE`
- [](#VUID-VkSubpassDescriptionDepthStencilResolve-pNext-06875)VUID-VkSubpassDescriptionDepthStencilResolve-pNext-06875  
  If the `pNext` chain of [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription2.html) includes a `VkMultisampledRenderToSingleSampledInfoEXT` structure whose `multisampledRenderToSingleSampledEnable` field is `VK_TRUE`, and `pDepthStencilAttachment` is not `NULL`, does not have the value `VK_ATTACHMENT_UNUSED`, and has a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) with a stencil component, then the value of `stencilResolveMode` **must** be one of the bits set in [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`supportedStencilResolveModes` or `VK_RESOLVE_MODE_NONE`
- [](#VUID-VkSubpassDescriptionDepthStencilResolve-pNext-06876)VUID-VkSubpassDescriptionDepthStencilResolve-pNext-06876  
  If the `pNext` chain of [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription2.html) includes a `VkMultisampledRenderToSingleSampledInfoEXT` structure whose `multisampledRenderToSingleSampledEnable` field is `VK_TRUE`, `pDepthStencilAttachment` is not `NULL`, does not have the value `VK_ATTACHMENT_UNUSED`, and has a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) with both depth and stencil components, and both [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`independentResolve` and [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`independentResolveNone` are `VK_FALSE`, then the values of `depthResolveMode` and `stencilResolveMode` **must** be identical
- [](#VUID-VkSubpassDescriptionDepthStencilResolve-pNext-06877)VUID-VkSubpassDescriptionDepthStencilResolve-pNext-06877  
  If the `pNext` chain of [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription2.html) includes a `VkMultisampledRenderToSingleSampledInfoEXT` structure whose `multisampledRenderToSingleSampledEnable` field is `VK_TRUE`, `pDepthStencilAttachment` is not `NULL`, does not have the value `VK_ATTACHMENT_UNUSED`, and has a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) with both depth and stencil components, [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`independentResolve` is `VK_FALSE`, and [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`independentResolveNone` is `VK_TRUE`, then the values of `depthResolveMode` and `stencilResolveMode` **must** be identical or one of them **must** be `VK_RESOLVE_MODE_NONE`

Valid Usage (Implicit)

- [](#VUID-VkSubpassDescriptionDepthStencilResolve-sType-sType)VUID-VkSubpassDescriptionDepthStencilResolve-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_DEPTH_STENCIL_RESOLVE`
- [](#VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-parameter)VUID-VkSubpassDescriptionDepthStencilResolve-pDepthStencilResolveAttachment-parameter  
  If `pDepthStencilResolveAttachment` is not `NULL`, `pDepthStencilResolveAttachment` **must** be a valid pointer to a valid [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference2.html) structure

## [](#_see_also)See Also

[VK\_KHR\_depth\_stencil\_resolve](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_depth_stencil_resolve.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference2.html), [VkResolveModeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResolveModeFlagBits.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSubpassDescriptionDepthStencilResolve).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0