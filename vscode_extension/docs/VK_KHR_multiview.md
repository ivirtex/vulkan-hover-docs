# VK\_KHR\_multiview(3) Manual Page

## Name

VK\_KHR\_multiview - device extension



## [](#_registered_extension_number)Registered Extension Number

54

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_multiview](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_multiview.html)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.1](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.1-promotions)

## [](#_contact)Contact

- Jeff Bolz [\[GitHub\]jeffbolznv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_multiview%5D%20%40jeffbolznv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_multiview%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2016-10-28

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension provides API support for [`GL_EXT_multiview`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GL_EXT_multiview.txt)

**Contributors**

- Jeff Bolz, NVIDIA

## [](#_description)Description

This extension has the same goal as the OpenGL ES `GL_OVR_multiview` extension. Multiview is a rendering technique originally designed for VR where it is more efficient to record a single set of commands to be executed with slightly different behavior for each “view”.

It includes a concise way to declare a render pass with multiple views, and gives implementations freedom to render the views in the most efficient way possible. This is done with a multiview configuration specified during [render pass](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass) creation with the [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassMultiviewCreateInfo.html) passed into [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo.html)::`pNext`.

This extension enables the use of the [`SPV_KHR_multiview`](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_multiview.html) shader extension, which adds a new `ViewIndex` built-in type that allows shaders to control what to do for each view. If using GLSL there is also the [`GL_EXT_multiview`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GL_EXT_multiview.txt) extension that introduces a `highp int gl_ViewIndex;` built-in variable for vertex, tessellation, geometry, and fragment shaders.

## [](#_promotion_to_vulkan_1_1)Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with the KHR suffix omitted. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceMultiviewFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMultiviewFeaturesKHR.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceMultiviewPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMultiviewPropertiesKHR.html)
- Extending [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo.html):
  
  - [VkRenderPassMultiviewCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassMultiviewCreateInfoKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_MULTIVIEW_EXTENSION_NAME`
- `VK_KHR_MULTIVIEW_SPEC_VERSION`
- Extending [VkDependencyFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyFlagBits.html):
  
  - `VK_DEPENDENCY_VIEW_LOCAL_BIT_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PROPERTIES_KHR`
  - `VK_STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO_KHR`

## [](#_new_built_in_variables)New Built-In Variables

- [`ViewIndex`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-viewindex)

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`MultiView`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-MultiView)

## [](#_additional_resources)Additional Resources

- ['NVIDIA blog post'](https://devblogs.nvidia.com/turing-multi-view-rendering-vrworks)
- ['ARM blog post'](https://community.arm.com/developer/tools-software/graphics/b/blog/posts/optimizing-virtual-reality-understanding-multiview)

## [](#_version_history)Version History

- Revision 1, 2016-10-28 (Jeff Bolz)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_multiview)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0