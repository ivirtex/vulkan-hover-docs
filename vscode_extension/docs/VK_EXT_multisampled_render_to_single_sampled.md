# VK\_EXT\_multisampled\_render\_to\_single\_sampled(3) Manual Page

## Name

VK\_EXT\_multisampled\_render\_to\_single\_sampled - device extension



## [](#_registered_extension_number)Registered Extension Number

377

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_create\_renderpass2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_create_renderpass2.html)  
     and  
     [VK\_KHR\_depth\_stencil\_resolve](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_depth_stencil_resolve.html)  
or  
[Vulkan Version 1.2](#versions-1.2)

## [](#_contact)Contact

- Shahbaz Youssefi [\[GitHub\]syoussefi](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_multisampled_render_to_single_sampled%5D%20%40syoussefi%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_multisampled_render_to_single_sampled%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_multisampled\_render\_to\_single\_sampled](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_multisampled_render_to_single_sampled.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-04-16

**IP Status**

No known IP claims.

**Contributors**

- Shahbaz Youssefi, Google
- Jan-Harald Fredriksen, Arm
- Jörg Wagner, Arm
- Matthew Netsch, Qualcomm Technologies, Inc.
- Jarred Davies, Imagination Technologies

## [](#_description)Description

With careful usage of resolve attachments, multisampled image memory allocated with `VK_MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT`, `loadOp` not equal to `VK_ATTACHMENT_LOAD_OP_LOAD` and `storeOp` not equal to `VK_ATTACHMENT_STORE_OP_STORE`, a Vulkan application is able to efficiently perform multisampled rendering without incurring any additional memory penalty on some implementations.

Under certain circumstances however, the application may not be able to complete its multisampled rendering within a single render pass; for example if it does partial rasterization from frame to frame, blending on an image from a previous frame, or in emulation of GL\_EXT\_multisampled\_render\_to\_texture. In such cases, the application can use an initial subpass to effectively load single-sampled data from the next subpass’s resolve attachment and fill in the multisampled attachment which otherwise uses `loadOp` equal to `VK_ATTACHMENT_LOAD_OP_DONT_CARE`. However, this is not always possible (for example for stencil in the absence of VK\_EXT\_shader\_stencil\_export) and has multiple drawbacks.

Some implementations are able to perform said operation efficiently in hardware, effectively loading a multisampled attachment from the contents of a single sampled one. Together with the ability to perform a resolve operation at the end of a subpass, these implementations are able to perform multisampled rendering on single-sampled attachments with no extra memory or bandwidth overhead. This extension exposes this capability by allowing a framebuffer and render pass to include single-sampled attachments while rendering is done with a specified number of samples.

## [](#_new_structures)New Structures

- Extending [VkFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties2.html):
  
  - [VkSubpassResolvePerformanceQueryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassResolvePerformanceQueryEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT.html)
- Extending [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription2.html), [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html):
  
  - [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_MULTISAMPLED_RENDER_TO_SINGLE_SAMPLED_EXTENSION_NAME`
- `VK_EXT_MULTISAMPLED_RENDER_TO_SINGLE_SAMPLED_SPEC_VERSION`
- Extending [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateFlagBits.html):
  
  - `VK_IMAGE_CREATE_MULTISAMPLED_RENDER_TO_SINGLE_SAMPLED_BIT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_MULTISAMPLED_RENDER_TO_SINGLE_SAMPLED_INFO_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTISAMPLED_RENDER_TO_SINGLE_SAMPLED_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_SUBPASS_RESOLVE_PERFORMANCE_QUERY_EXT`

## [](#_issues)Issues

1\) Could the multisampled attachment be initialized through some form of copy?

**RESOLVED**: No. Some implementations do not support copying between attachments in general, and find expressing this operation through a copy unnatural.

2\) Another way to achieve this is by introducing a new `loadOp` to load the contents of the multisampled image from a single-sampled one. Why is this extension preferred?

**RESOLVED**: Using this extension simplifies the application, as it does not need to manage a secondary lazily-allocated image. Additionally, using this extension leaves less room for error; for example a single mistake in `loadOp` or `storeOp` would result in the lazily-allocated image to actually take up memory, and remain so until destruction.

3\) There is no guarantee that multisampled data between two subpasses with the same number of samples will be retained as the implementation may be forced to split the render pass implicitly for various reasons. Should this extension require that every subpass that uses multisampled-render-to-single-sampled end in an implicit render pass split (which results in a resolve operation)?

**RESOLVED**: No. Not requiring this allows render passes with multiple multisampled-render-to-single-sampled subpasses to potentially execute more efficiently (though there is no guarantee).

## [](#_version_history)Version History

- Revision 1, 2021-04-12 (Shahbaz Youssefi)

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_multisampled_render_to_single_sampled).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0